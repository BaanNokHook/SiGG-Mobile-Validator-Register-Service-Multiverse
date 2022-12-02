[Authorize("Users")]
 public class ProfileController : ResourceController
    {
        private readonly UserReports _userReports;
        private readonly Func<UserService> _userServiceFactory;
        private readonly ILogger _logger;
        private readonly Func<ProfileViewModelValidator> _validatorFactory;

        public ProfileController(UserReports userReports, Func<UserService> userServiceFactory, Func<ProfileViewModelValidator> validatorFactory, ILogger logger)
        {
            _userReports = userReports;
            _logger = logger;
            _userServiceFactory = userServiceFactory;
            _validatorFactory = validatorFactory;
        }

        [Get("/accounts/profile")]
        public ActionResult Get()
        {            
            var user = _userReports.FindBy(User.Identity.Name);

            if (user == null)
            {  
                _logger.Write("Could not find user for " + User.Identity.Name);
                return new NotFound();
            }

            var profile = DomainToPublic.Map(user, new ProfileViewModel());
            
            return new OK(profile);
        }

        [Put("/accounts/profile")]
        public ActionResult Put(ProfileViewModel profileViewModel)
        {
            var validationResult = _validatorFactory().Validate(profileViewModel);
            
            var user = _userReports.FindBy(User.Identity.Name);

            if (validationResult.IsValid)
            {
                _userServiceFactory().Update(user.Id, user.ETag, (dbUser) => PublicToDomain.Map(profileViewModel, dbUser));
                return new OkSeeOther("?profileUpdated=true");
            }
            
            return new ForbiddenFor(string.Join(Environment.NewLine, validationResult.Errors.Select(x => x.ErrorMessage)));
        }

        [Authorize("Administrators")]
        [Delete("/accounts/profile")]
        public ActionResult Delete()
        {
           ...
        }
    }