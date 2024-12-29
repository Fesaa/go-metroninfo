package metroninfo

type Format string

const (
	AnnualFormat         Format = "Annual"
	DigitalChapterFormat Format = "Digital Chapter"
	GraphicNovelFormat   Format = "Graphic Novel"
	HardcoverFormat      Format = "Hardcover"
	LimitedSeriesFormat  Format = "Limited Series" // Used for mini/maxi series
	OmnibusFormat        Format = "Omnibus"
	OneShotFormat        Format = "One-Shot"
	SingleIssueFormat    Format = "Single Issue" // Used for Ongoing/Cancelled series
	TradePaperbackFormat Format = "Trade Paperback"
)

type InformationSource string

const (
	SourceAniList             InformationSource = "AniList"
	SourceComicVine           InformationSource = "Comic Vine"
	SourceGrandComicsDatabase InformationSource = "Grand Comics Database"
	SourceKitsu               InformationSource = "Kitsu"    // IDs contain letters, hyphens, and numbers
	SourceMangaDex            InformationSource = "MangaDex" // IDs contain letters, hyphens, and numbers
	SourceMangaUpdates        InformationSource = "MangaUpdates"
	SourceMarvel              InformationSource = "Marvel"
	SourceMetron              InformationSource = "Metron"
	SourceMyAnimeList         InformationSource = "MyAnimeList"
	SourceLeagueOfComicGeeks  InformationSource = "League of Comic Geeks"
)

type RoleValue string

const (
	RoleWriter               RoleValue = "Writer"
	RoleScript               RoleValue = "Script"
	RoleStory                RoleValue = "Story"
	RolePlot                 RoleValue = "Plot"
	RoleInterviewer          RoleValue = "Interviewer"
	RoleArtist               RoleValue = "Artist"
	RolePenciller            RoleValue = "Penciller"
	RoleBreakdowns           RoleValue = "Breakdowns"
	RoleIllustrator          RoleValue = "Illustrator"
	RoleLayouts              RoleValue = "Layouts"
	RoleInker                RoleValue = "Inker"
	RoleEmbellisher          RoleValue = "Embellisher"
	RoleFinishes             RoleValue = "Finishes"
	RoleInkAssists           RoleValue = "Ink Assists"
	RoleColorist             RoleValue = "Colorist"
	RoleColorSeparations     RoleValue = "Color Separations"
	RoleColorAssists         RoleValue = "Color Assists"
	RoleColorFlats           RoleValue = "Color Flats"
	RoleDigitalArtTechnician RoleValue = "Digital Art Technician"
	RoleGrayTone             RoleValue = "Gray Tone"
	RoleLetterer             RoleValue = "Letterer"
	RoleCover                RoleValue = "Cover"
	RoleEditor               RoleValue = "Editor"
	RoleConsultingEditor     RoleValue = "Consulting Editor"
	RoleAssistantEditor      RoleValue = "Assistant Editor"
	RoleAssociateEditor      RoleValue = "Associate Editor"
	RoleGroupEditor          RoleValue = "Group Editor"
	RoleSeniorEditor         RoleValue = "Senior Editor"
	RoleManagingEditor       RoleValue = "Managing Editor"
	RoleCollectionEditor     RoleValue = "Collection Editor"
	RoleProduction           RoleValue = "Production"
	RoleDesigner             RoleValue = "Designer"
	RoleLogoDesign           RoleValue = "Logo Design"
	RoleTranslator           RoleValue = "Translator"
	RoleSupervisingEditor    RoleValue = "Supervising Editor"
	RoleExecutiveEditor      RoleValue = "Executive Editor"
	RoleEditorInChief        RoleValue = "Editor In Chief"
	RolePresident            RoleValue = "President"
	RolePublisher            RoleValue = "Publisher"
	RoleChiefCreativeOfficer RoleValue = "Chief Creative Officer"
	RoleExecutiveProducer    RoleValue = "Executive Producer"
	RoleOther                RoleValue = "Other"
)

type AgeRating string

const (
	AgeRatingUnknown  AgeRating = "Unknown"
	AgeRatingEveryone AgeRating = "Everyone"  // Appropriate for readers of all ages.
	AgeRatingTeen     AgeRating = "Teen"      // Appropriate for readers age 12 and older.
	AgeRatingTeenPlus AgeRating = "Teen Plus" // Appropriate for readers age 15 and older.
	AgeRatingMature   AgeRating = "Mature"    // Appropriate for readers age 17 and older.
	AgeRatingExplicit AgeRating = "Explicit"  // Contains Gore, Sexually Explicit material that would be more extreme than R rating.
	AgeRatingAdult    AgeRating = "Adult"     // Likely pornographic in nature.
)
