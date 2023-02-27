package seeders

import (
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) CountryCode() error {
	listOfObject := []models.CountryCode{
		{Id: 1, Code: "India", Description: "101"},
		{Id: 2, Code: "Albania", Description: "003"},
		{Id: 3, Code: "Aland Islands", Description: "002"},
		{Id: 4, Code: "Afghanistan", Description: "001"},
		{Id: 5, Code: "Algeria", Description: "004"},
		{Id: 6, Code: "American Samoa", Description: "005"},
		{Id: 7, Code: "Andorra", Description: "006"},
		{Id: 8, Code: "Angola", Description: "007"},
		{Id: 9, Code: "Anguilla", Description: "008"},
		{Id: 10, Code: "Antarctica", Description: "009"},
		{Id: 11, Code: "Antigua And Barbuda", Description: "010"},
		{Id: 12, Code: "Argentina", Description: "011"},
		{Id: 13, Code: "Armenia", Description: "012"},
		{Id: 14, Code: "Aruba", Description: "013"},
		{Id: 15, Code: "Australia", Description: "014"},
		{Id: 16, Code: "Austria", Description: "015"},
		{Id: 17, Code: "Azerbaijan", Description: "016"},
		{Id: 18, Code: "Bahamas", Description: "017"},
		{Id: 19, Code: "Bahrain", Description: "018"},
		{Id: 20, Code: "Bangladesh", Description: "019"},
		{Id: 21, Code: "Barbados", Description: "020"},
		{Id: 22, Code: "Belarus", Description: "021"},
		{Id: 23, Code: "Belgium", Description: "022"},
		{Id: 24, Code: "Belize", Description: "023"},
		{Id: 25, Code: "Benin", Description: "024"},
		{Id: 26, Code: "Bermuda", Description: "025"},
		{Id: 27, Code: "Bhutan", Description: "026"},
		{Id: 28, Code: "Bolivia", Description: "027"},
		{Id: 29, Code: "Bosnia And Herzegovina", Description: "028"},
		{Id: 30, Code: "Botswana", Description: "029"},
		{Id: 31, Code: "Bouvet Island", Description: "030"},
		{Id: 32, Code: "Brazil", Description: "031"},
		{Id: 33, Code: "British Indian Ocean Territory", Description: "032"},
		{Id: 34, Code: "Brunei Darussalam", Description: "033"},
		{Id: 35, Code: "Bulgaria", Description: "034"},
		{Id: 36, Code: "Burkina Faso", Description: "035"},
		{Id: 37, Code: "Burundi", Description: "036"},
		{Id: 38, Code: "Cambodia", Description: "037"},
		{Id: 39, Code: "Cameroon", Description: "038"},
		{Id: 40, Code: "Canada", Description: "039"},
		{Id: 41, Code: "Cape Verde", Description: "040"},
		{Id: 42, Code: "Cayman Islands", Description: "041"},
		{Id: 43, Code: "Central African Republic", Description: "042"},
		{Id: 44, Code: "Chad", Description: "043"},
		{Id: 45, Code: "Chile", Description: "044"},
		{Id: 46, Code: "China", Description: "045"},
		{Id: 47, Code: "Christmas Island", Description: "046"},
		{Id: 48, Code: "Cocos (Keeling) Islands", Description: "047"},
		{Id: 49, Code: "Colombia", Description: "048"},
		{Id: 50, Code: "Comoros", Description: "049"},
		{Id: 51, Code: "Congo", Description: "050"},
		{Id: 52, Code: "Congo, The Democratic Republic Of The", Description: "051"},
		{Id: 53, Code: "Cook Islands", Description: "052"},
		{Id: 54, Code: "Costa Rica", Description: "053"},
		{Id: 55, Code: "Cote D'Ivoire", Description: "054"},
		{Id: 56, Code: "Croatia", Description: "055"},
		{Id: 57, Code: "Cuba", Description: "056"},
		{Id: 58, Code: "Cyprus", Description: "057"},
		{Id: 59, Code: "Czech Republic", Description: "058"},
		{Id: 60, Code: "Denmark", Description: "059"},
		{Id: 61, Code: "Djibouti", Description: "060"},
		{Id: 62, Code: "Dominica", Description: "061"},
		{Id: 63, Code: "Dominican Republic", Description: "062"},
		{Id: 64, Code: "Ecuador", Description: "063"},
		{Id: 65, Code: "Egypt", Description: "064"},
		{Id: 66, Code: "El Salvador", Description: "065"},
		{Id: 67, Code: "Equatorial Guinea", Description: "066"},
		{Id: 68, Code: "Eritrea", Description: "067"},
		{Id: 69, Code: "Estonia", Description: "068"},
		{Id: 70, Code: "Ethiopia", Description: "069"},
		{Id: 71, Code: "Falkland Islands (Malvinas)", Description: "070"},
		{Id: 72, Code: "Faroe Islands", Description: "071"},
		{Id: 73, Code: "Fiji", Description: "072"},
		{Id: 74, Code: "Finland", Description: "073"},
		{Id: 75, Code: "France", Description: "074"},
		{Id: 76, Code: "French Guiana", Description: "075"},
		{Id: 77, Code: "French Polynesia", Description: "076"},
		{Id: 78, Code: "French Southern Territories", Description: "077"},
		{Id: 79, Code: "Gabon", Description: "078"},
		{Id: 80, Code: "Gambia", Description: "079"},
		{Id: 81, Code: "Georgia", Description: "080"},
		{Id: 82, Code: "Germany", Description: "081"},
		{Id: 83, Code: "Ghana", Description: "082"},
		{Id: 84, Code: "Gibraltar", Description: "083"},
		{Id: 85, Code: "Greece", Description: "084"},
		{Id: 86, Code: "Greenland", Description: "085"},
		{Id: 87, Code: "Grenada", Description: "086"},
		{Id: 88, Code: "Guadeloupe", Description: "087"},
		{Id: 89, Code: "Guam", Description: "088"},
		{Id: 90, Code: "Guatemala", Description: "089"},
		{Id: 91, Code: "Guernsey", Description: "090"},
		{Id: 92, Code: "Guinea", Description: "091"},
		{Id: 93, Code: "Guinea-Bissau", Description: "092"},
		{Id: 94, Code: "Guyana", Description: "093"},
		{Id: 95, Code: "Haiti", Description: "094"},
		{Id: 96, Code: "Heard Island And Mcdonald Islands", Description: "095"},
		{Id: 97, Code: "Holy See (Vatican City State)", Description: "096"},
		{Id: 98, Code: "Honduras", Description: "097"},
		{Id: 99, Code: "Hong Kong", Description: "098"},
		{Id: 100, Code: "Hungary", Description: "099"},
		{Id: 101, Code: "Iceland", Description: "100"},
		{Id: 102, Code: "Indonesia", Description: "102"},
		{Id: 103, Code: "Iran, Islamic Republic Of", Description: "103"},
		{Id: 104, Code: "Iraq", Description: "104"},
		{Id: 105, Code: "Ireland", Description: "105"},
		{Id: 106, Code: "Isle Of Man", Description: "106"},
		{Id: 107, Code: "Israel", Description: "107"},
		{Id: 108, Code: "Italy", Description: "108"},
		{Id: 109, Code: "Jamaica", Description: "109"},
		{Id: 110, Code: "Japan", Description: "110"},
		{Id: 111, Code: "Jersey", Description: "111"},
		{Id: 112, Code: "Jordan", Description: "112"},
		{Id: 113, Code: "Kazakhstan", Description: "113"},
		{Id: 114, Code: "Kenya", Description: "114"},
		{Id: 115, Code: "Kiribati", Description: "115"},
		{Id: 116, Code: "Korea, Democratic People’s Republic Of", Description: "116"},
		{Id: 117, Code: "Korea, Republic Of", Description: "117"},
		{Id: 118, Code: "Kuwait", Description: "118"},
		{Id: 119, Code: "Kyrgyzstan", Description: "119"},
		{Id: 120, Code: "Lao People’s Democratic Republic", Description: "120"},
		{Id: 121, Code: "Latvia", Description: "121"},
		{Id: 122, Code: "Lebanon", Description: "122"},
		{Id: 123, Code: "Lesotho", Description: "123"},
		{Id: 124, Code: "Liberia", Description: "124"},
		{Id: 125, Code: "Libyan Arab Jamahiriya", Description: "125"},
		{Id: 126, Code: "Liechtenstein", Description: "126"},
		{Id: 127, Code: "Lithuania", Description: "127"},
		{Id: 128, Code: "Luxembourg", Description: "128"},
		{Id: 129, Code: "Macao", Description: "129"},
		{Id: 130, Code: "Macedonia, The Former Yugoslav Republic Of", Description: "130"},
		{Id: 131, Code: "Madagascar", Description: "131"},
		{Id: 132, Code: "Malawi", Description: "132"},
		{Id: 133, Code: "Malaysia", Description: "133"},
		{Id: 134, Code: "Maldives", Description: "134"},
		{Id: 135, Code: "Mali", Description: "135"},
		{Id: 136, Code: "Malta", Description: "136"},
		{Id: 137, Code: "Marshall Islands", Description: "137"},
		{Id: 138, Code: "Martinique", Description: "138"},
		{Id: 139, Code: "Mauritania", Description: "139"},
		{Id: 140, Code: "Mauritius", Description: "140"},
		{Id: 141, Code: "Mayotte", Description: "141"},
		{Id: 142, Code: "Mexico", Description: "142"},
		{Id: 143, Code: "Micronesia, Federated States Of", Description: "143"},
		{Id: 144, Code: "Moldova, Republic Of", Description: "144"},
		{Id: 145, Code: "Monaco", Description: "145"},
		{Id: 146, Code: "Mongolia", Description: "146"},
		{Id: 147, Code: "Montserrat", Description: "147"},
		{Id: 148, Code: "Morocco", Description: "148"},
		{Id: 149, Code: "Mozambique", Description: "149"},
		{Id: 150, Code: "Myanmar", Description: "150"},
		{Id: 151, Code: "Namibia", Description: "151"},
		{Id: 152, Code: "Nauru", Description: "152"},
		{Id: 153, Code: "Nepal", Description: "153"},
		{Id: 154, Code: "Netherlands", Description: "154"},
		{Id: 155, Code: "Netherlands Antilles", Description: "155"},
		{Id: 156, Code: "New Caledonia", Description: "156"},
		{Id: 157, Code: "New Zealand", Description: "157"},
		{Id: 158, Code: "Nicaragua", Description: "158"},
		{Id: 159, Code: "Niger", Description: "159"},
		{Id: 160, Code: "Nigeria", Description: "160"},
		{Id: 161, Code: "Niue", Description: "161"},
		{Id: 162, Code: "Norfolk Island", Description: "162"},
		{Id: 163, Code: "Northern Mariana Islands", Description: "163"},
		{Id: 164, Code: "Norway", Description: "164"},
		{Id: 165, Code: "Oman", Description: "165"},
		{Id: 166, Code: "Pakistan", Description: "166"},
		{Id: 167, Code: "Palau", Description: "167"},
		{Id: 168, Code: "Palestinian Territory, Occupied", Description: "168"},
		{Id: 169, Code: "Panama", Description: "169"},
		{Id: 170, Code: "Papua New Guinea", Description: "170"},
		{Id: 171, Code: "Paraguay", Description: "171"},
		{Id: 172, Code: "Peru", Description: "172"},
		{Id: 173, Code: "Philippines", Description: "173"},
		{Id: 174, Code: "Pitcairn", Description: "174"},
		{Id: 175, Code: "Poland", Description: "175"},
		{Id: 176, Code: "Portugal", Description: "176"},
		{Id: 177, Code: "Puerto Rico", Description: "177"},
		{Id: 178, Code: "Qatar", Description: "178"},
		{Id: 179, Code: "Reunion", Description: "179"},
		{Id: 180, Code: "Romania", Description: "180"},
		{Id: 181, Code: "Russian Federation", Description: "181"},
		{Id: 182, Code: "Rwanda", Description: "182"},
		{Id: 183, Code: "Saint Helena", Description: "183"},
		{Id: 184, Code: "Saint Kitts And Nevis", Description: "184"},
		{Id: 185, Code: "Saint Lucia", Description: "185"},
		{Id: 186, Code: "Saint Pierre And Miquelon", Description: "186"},
		{Id: 187, Code: "Saint Vincent And The Grenadines", Description: "187"},
		{Id: 188, Code: "Samoa", Description: "188"},
		{Id: 189, Code: "San Marino", Description: "189"},
		{Id: 190, Code: "Sao Tome And Principe", Description: "190"},
		{Id: 191, Code: "Saudi Arabia", Description: "191"},
		{Id: 192, Code: "Senegal", Description: "192"},
		{Id: 193, Code: "Serbia And Montenegro", Description: "193"},
		{Id: 194, Code: "Seychelles", Description: "194"},
		{Id: 195, Code: "Sierra Leone", Description: "195"},
		{Id: 196, Code: "Singapore", Description: "196"},
		{Id: 197, Code: "Slovakia", Description: "197"},
		{Id: 198, Code: "Slovenia", Description: "198"},
		{Id: 199, Code: "Solomon Islands", Description: "199"},
		{Id: 200, Code: "Somalia", Description: "200"},
		{Id: 201, Code: "South Africa", Description: "201"},
		{Id: 202, Code: "South Georgia And The South Sandwich Islands", Description: "202"},
		{Id: 203, Code: "Spain", Description: "203"},
		{Id: 204, Code: "Sri Lanka", Description: "204"},
		{Id: 205, Code: "Sudan", Description: "205"},
		{Id: 206, Code: "Suriname", Description: "206"},
		{Id: 207, Code: "Svalbard And Jan Mayen", Description: "207"},
		{Id: 208, Code: "Swaziland", Description: "208"},
		{Id: 209, Code: "Sweden", Description: "209"},
		{Id: 210, Code: "Switzerland", Description: "210"},
		{Id: 211, Code: "Syrian Arab Republic", Description: "211"},
		{Id: 212, Code: "Taiwan, Province Of China", Description: "212"},
		{Id: 213, Code: "Tajikistan", Description: "213"},
		{Id: 214, Code: "Tanzania, United Republic Of", Description: "214"},
		{Id: 215, Code: "Thailand", Description: "215"},
		{Id: 216, Code: "Timor-Leste", Description: "216"},
		{Id: 217, Code: "Togo", Description: "217"},
		{Id: 218, Code: "Tokelau", Description: "218"},
		{Id: 219, Code: "Tonga", Description: "219"},
		{Id: 220, Code: "Trinidad And Tobago", Description: "220"},
		{Id: 221, Code: "Tunisia", Description: "221"},
		{Id: 222, Code: "Turkey", Description: "222"},
		{Id: 223, Code: "Turkmenistan", Description: "223"},
		{Id: 224, Code: "Turks And Caicos Islands", Description: "224"},
		{Id: 225, Code: "Tuvalu", Description: "225"},
		{Id: 226, Code: "Uganda", Description: "226"},
		{Id: 227, Code: "Ukraine", Description: "227"},
		{Id: 228, Code: "United Arab Emirates", Description: "228"},
		{Id: 229, Code: "United Kingdom", Description: "229"},
		{Id: 230, Code: "United States", Description: "230"},
		{Id: 231, Code: "United States Minor Outlying Islands", Description: "231"},
		{Id: 232, Code: "Uruguay", Description: "232"},
		{Id: 233, Code: "Uzbekistan", Description: "233"},
		{Id: 234, Code: "Vanuatu", Description: "234"},
		{Id: 235, Code: "Venezuela", Description: "235"},
		{Id: 236, Code: "Viet Nam", Description: "236"},
		{Id: 237, Code: "Virgin Islands, British", Description: "237"},
		{Id: 238, Code: "Virgin Islands, U.S.", Description: "238"},
		{Id: 239, Code: "Wallis And Futuna", Description: "239"},
		{Id: 240, Code: "Western Sahara", Description: "240"},
		{Id: 241, Code: "Yemen", Description: "241"},
		{Id: 242, Code: "Zambia", Description: "242"},
		{Id: 243, Code: "Zimbabwe", Description: "243"},
		{Id: 244, Code: "Côte D'ivoire", Description: "CI"},
		{Id: 245, Code: "Korea,Democratic People'sRepublicOf", Description: "KP"},
		{Id: 246, Code: "Lao People’s Democratic Republic", Description: "12"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.FirstOrCreate(&listDb)
		if resp.Error != nil {
			utils.Log.Error(resp.Error)
		}
	}
	return nil
}
