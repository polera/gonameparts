package gonameparts

var (
	salutations      = []string{"MR", "MS", "MRS", "DR", "MISS", "DOCTOR", "CORP", "SGT", "PVT", "JUDGE", "CAPT", "COL", "MAJ", "LT", "LIEUTENANT", "PRM", "PATROLMAN", "HON", "OFFICER", "REV", "PRES", "PRESIDENT", "GOV", "GOVERNOR", "VICE PRESIDENT", "VP", "MAYOR", "SIR", "MADAM", "HONERABLE"}
	generations      = []string{"JR", "SR", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "1ST", "2ND", "3RD", "4TH", "5TH", "6TH", "7TH", "8TH", "9TH", "10TH", "FIRST", "SECOND", "THIRD", "FOURTH", "FIFTH", "SIXTH", "SEVENTH", "EIGHTH", "NINTH", "TENTH"}
	suffixes         = []string{"ESQ", "PHD", "MD"}
	lnPrefixes       = []string{"DE", "DA", "DI", "LA", "DU", "DEL", "DEI", "VDA", "DELLO", "DELLA", "DEGLI", "DELLE", "VAN", "VON", "DER", "DEN", "HEER", "TEN", "TER", "VANDE", "VANDEN", "VANDER", "VOOR", "VER", "AAN", "MC", "BEN", "SAN", "SAINZ", "BIN", "LI", "LE", "DES", "AM", "AUS'M", "VOM", "ZUM", "ZUR", "TEN", "IBN"}
	nonName          = []string{"A.K.A", "AKA", "A/K/A", "F.K.A", "FKA", "F/K/A", "N/K/A"}
	corpEntity       = []string{"NA", "CORP", "CO", "INC", "ASSOCIATES", "SERVICE", "LLC", "LLP", "PARTNERS", "R/A", "C/O", "COUNTY", "STATE", "BANK", "GROUP", "MUTUAL", "FARGO"}
	supplementalInfo = []string{"WIFE OF", "HUSBAND OF", "SON OF", "DAUGHTER OF", "DECEASED", "FICTITIOUS"}
)
