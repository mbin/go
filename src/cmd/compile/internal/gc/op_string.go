// Code generated by "stringer -type=Op -trimprefix=O"; DO NOT EDIT.

package gc

import "strconv"

const _Op_name = "XXXNAMENONAMETYPEPACKLITERALADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESASAS2AS2FUNCAS2RECVAS2MAPRAS2DOTTYPEASOPCALLCALLFUNCCALLMETHCALLINTERCALLPARTCAPCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVNOPCOPYDCLDCLFUNCDCLFIELDDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMULDIVMODLSHRSHANDANDNOTNEWNEWOBJNOTBITNOTPLUSNEGORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERRECOVERRECVRUNESTRSELRECVSELRECV2IOTAREALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFBLOCKBREAKCASEXCASECONTINUEDEFEREMPTYFALLFORFORUNTILGOTOIFLABELGORANGERETURNSELECTSWITCHTYPESWTCHANTMAPTSTRUCTTINTERTFUNCTARRAYDDDDDDARGINLCALLEFACEITABIDATASPTRCLOSUREVARCFUNCCHECKNILVARDEFVARKILLVARLIVEINDREGSPINLMARKRETJMPGETGEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 21, 28, 31, 34, 36, 39, 45, 49, 55, 61, 70, 82, 91, 100, 112, 121, 123, 126, 133, 140, 147, 157, 161, 165, 173, 181, 190, 198, 201, 206, 213, 220, 226, 235, 243, 251, 257, 261, 270, 277, 281, 284, 291, 299, 307, 314, 320, 323, 329, 336, 344, 348, 355, 363, 365, 367, 369, 371, 373, 375, 380, 385, 393, 396, 405, 408, 412, 420, 427, 436, 439, 442, 445, 448, 451, 454, 460, 463, 469, 472, 478, 482, 485, 489, 494, 499, 505, 510, 514, 519, 527, 535, 541, 550, 561, 568, 572, 579, 586, 594, 598, 602, 606, 613, 620, 628, 634, 639, 644, 648, 653, 661, 666, 671, 675, 678, 686, 690, 692, 697, 699, 704, 710, 716, 722, 728, 733, 737, 744, 750, 755, 761, 764, 770, 777, 782, 786, 791, 795, 805, 810, 818, 824, 831, 838, 846, 853, 859, 863, 866}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
