package daythree

import (
	"strings"
)

func DayThreePartOne(array []string) int {
	myString := ""
	count := 0
	count2 := 0
	for _, v := range array {
		count += 1
		firstHalf := v[:(len(v) / 2)]
		secondHalf := v[(len(v) / 2):]
		for _, v := range firstHalf {
			if strings.Contains(secondHalf, string(v)) {
				count2 += 1
				myString += string(v)
				break
			}
		}

	}
	return countString(myString)
}
func countString(myString string) int {
	smallLetters := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	uppercaseLetters := map[string]int{
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
	count := 0
	for _, v := range myString {
		match := smallLetters[string(v)]
		match2 := uppercaseLetters[string(v)]
		count += (match + match2)
	}
	return count
}

func PuzzleInput() []string {
	array := []string{
		"rGnRRccfcCRFDPqNWdwWJWRBdB",
		"jZzVVSZSjmQvZTSZfjmLzNPJqWtJBwqpNtBpdWdNvd",
		"fZfjlMLVshMFhcHnDG",
		"vJRmRfJbfRfJsCpvgLggNVsv",
		"zlzSrMPZMgBFFMNMVWsjsF",
		"dzBSBlzdcggRGdndnn",
		"hNwqVDVDdmQwQPrbDMSSMRSWSM",
		"LvnzJJtlcHstlffCVpMSbRMpBMbCgVWM",
		"lGvvscLHcfsHtvlsnZmmNGhNFVQqqTdqFd",
		"sHGzGsfcZnHfhbLRFrdrhRFf",
		"vwVqzSSjSSttjSqgNMqwzSSVCCgRPhPhFLdCRLLrdCRRrDLb",
		"vvwwtvvVwSvMQzjvmNtJsHBnsZllmnnTBBcBlTnW",
		"pzNVBVplhfLfVfVStZrZHbQHZQTb",
		"sdPPJdCjbdQHMTWt",
		"ngmJGFjJwJCmGcnJcFtgwcDBNLlqfBpvfBpqllhlfL",
		"TGpphMZhJQpGLZTMCCtsBMCSDsStBFcB",
		"jdfgHClHrdrbWvWgjqmctSDqltDsFFDFBc",
		"dNWWvfgCLzzNNLZz",
		"vvHzBrTSvHFbqZqZTBBtzdVfGCGhJSfGJGDSdGMhVG",
		"lgmgslssslscsnRlVGWGWGVMRCll",
		"pQPsjwnNgQmnNNwQgNLNnmgZqvFqtwqtrMMzvzwFtrqTrr",
		"TNhNLTrswwsnFNrrGZVnJnZmVpSjjVDnPp",
		"dBlzzCBqgfWMBpPSJDHVZSVP",
		"cqzfpcdcNFFGrhcb",
		"LLpzCzSzJnLQVnTNhQ",
		"vtDqwBcRztQbthtV",
		"qZZZqZvvsZwwjvjDjspFlPpSSCslpzPlls",
		"rsZVlmStlFZllvmSSvLRqcGMswMMzjMsjqGCjzMG",
		"VTPgQBQdBCwgjCnJGC",
		"pBNTNHdbdWLrDVLbDmLZ",
		"CFmsmwnCRmjCjnCzJZQhhBGQBsMzbP",
		"WvltTWVDdNttdlDbbZzJvPJvBBhhzM",
		"gHDfTNlDDdgVNdglgfqtpcSFwccCmmwnRSwzzHHF",
		"hgmhWgCdzpjPCzFC",
		"NVfgffZblZQVNtFZPntPDrJjPt",
		"TwbQLLfMgdhswWGH",
		"JStSPHPJNrBCtBZMPtHTVfVwLcbbLTcSgfbVfn",
		"GdhhCplmdWbVddLV",
		"qpFvhQppFRlqmFGppRhqvvFzZtQMHZMzBNBCCzNDHHMPzB",
		"MvCMQmJCMDQjwMJCSJpQmDcnGBggfsgljGNslsNGjTHTNf",
		"HRPtRRzVWWFhrFdZhTGlTTLLggBLVgGGLf",
		"zZhdWZPWRdqttWWRPzRPmcpmpQqvvSCvwHwQmJqD",
		"NwrVVWrvwtQDRdRqWbqh",
		"jCCZjlJZZGclGPpCJlcCBhJshgqTsQgbDdshTbgLRh",
		"ppppflPZlpvfffFzHqFV",
		"VmTwGVwnwHnSnqGSVmBBwwmTZvzbCrsNWcsZsqFzCWsbWrNv",
		"JMPPgRDPDgghJggtzsMbZvsNCrWcNsbM",
		"DhgJllPPJcgJpptDDDPldpmTwjHBmSjBBwmBwHBdBmTQ",
		"vblvFHvrQTjqqhCpFwLnFL",
		"RDJRWsdRdgdWZMCMQSVppMJpVL",
		"DzsgRNNsftWdQmcvbPcfcBfj",
		"wTgbsTmgqSbzSlSvFb",
		"rNdDdZnRtJFlmVSHZz",
		"mfNjjmmdNWhWCsffwC",
		"LltBNFLHHcJcfNdNTwbbNsfW",
		"vQSMCzQSGdSdvjsTwmwfwWwV",
		"nRCSzMCQDqlgdlRJHJ",
		"vgqTGwnhpLFsVvrR",
		"ZZzTHlTlJNbcLfRtppFzfs",
		"PlWBWbjQjJgCDCMPTTGT",
		"wtBdWdDpjDdwScBtnsFsPmmnPbPHHPghhh",
		"fSSVGVzlTlqGfffLTQMGHsmFsbmPFbsPbhQZFhhs",
		"TzLzfqJTVRGvRMvqvLGzMqRfSvcDpNjcjdcwSDBDtBDvSCDj",
		"FrdCCzmqFrpDRTRHRLnQJJnr",
		"SGfNRWlBZlgtRltGbvnTVbJnSHvLHVbH",
		"NRWsfWBwwqwDMDpc",
		"vSfsmsdssdjSdZdSgRnmLRGNzgGnqwwB",
		"VrlThFPTPjHFDPRqGGnRnNLqqqHG",
		"lPtTDVCCDrjppfdSQbcJMsbftv",
		"RFFslstRRfSljtsRJPjRtRtngSHbmbbhGbHQQmgbmgGhGQ",
		"DqBZdWZvTdMBBTLDzqBhFnvgVhnHHFbnQVnGnb",
		"wcwMMMTzZwsjtCtjFN",
		"CsscSlwwZDscMNhhpZhRZHRM",
		"VjQvjbQvbhJVbvTbnQTnHHLrRFRqqFqFVdMHNHNH",
		"QbTjjbtJnvttntjgjJjPDwwsBsSsSlslhlwwPCBS",
		"RLjhhBBcNNBNhNjhpwScwSCTFwcFvMwvlc",
		"dJqPJqbqtmgmgdPrdgwvDTCSSlMSFDlFTzFg",
		"stCCqmbGrrbmPsqZhHhpjjNhfjBnnsnp",
		"LfDzcMMVsDLLzNscGhFFcQhlhhBRll",
		"ZCbSbwdmwPnPnPmHjdjWFWsWhdGjdH",
		"bZnPbvnvqCvJDrLJNssf",
		"bbldQHVltHQSSJJmTlwJGjCjCChchgMhHhprCpMv",
		"WNPDFfqfzFsBFFnmgjGhgccjpcpgMGrs",
		"mZFFzDnmzBBZltdlZSVtZS",
		"hdJZZdJTvnJmRphvvpGnwvqzVFSwVlHVlFHFHNlzlgFgHF",
		"QWWWBbrCCrtjQsMdcWMBfttHsgFHVHgFSNzDzFlVLggHzz",
		"jQdBPCtrtcPcMQCTnqnhTvPGGvnJGR",
		"ZhsmsbNVmFssmblMMpdvvQdwLQRppRvQ",
		"GjnNNGNCGjfJcqrQDvQQpvQLPpnQQw",
		"BJjWqWSCrNHrqjfhlHzlVbVtFmtHTV",
		"mtJmPmQBbPFshnJzZpLZ",
		"gMqMHRGrCHSvrRGrqPZLnplFzlVhFspL",
		"RrDrgMGRDPNtfmDN",
		"wqqvqVmlpqchqrDD",
		"gzRltSjgFlWshrdngrpPcDcd",
		"RRjtsbGfGLbsLWSLtjzGSHNNHmHBwNNHNmNfllwVfl",
		"mqFBQgnMQQbJqnTswSMNWGsDswZZ",
		"HgzlRzHccfHsfTwSWfNSNf",
		"VLldvpHLVHrFJFJgpnQbJn",
		"SzCJtLdDCCtqCcdDfZMVMfGVbsVZPPzz",
		"wpwWQWjQgwQgjNBwgHQgsGfPfbPsZPGSPjZmrPrV",
		"QFwpTwpHlnHFNFWQWlQgNwNLCddvDShtLJnSJCLqJJttDt",
		"HMgCVHggtqMMvjgNjNCMMwfWfPwWPJQQNzWwJzlQlQ",
		"bGFrqrFGZLLdFmSPcmPzQJlQcWzJmJ",
		"hRphhShnLGFdGSFSLRGrdqdqsjTtsBTVTTBnVjVvHHBtHggt",
		"hHhnFHQpcFcHjcjfZfZRnjjnNNBvNNNtwvNtbwPtNcPtBgbg",
		"zzVLWCHLSdCttbvw",
		"mMLsHDMVFlRhhmhm",
		"MHMgBNQQPCMMQBbBQQgJHbWttdlfWpZJVWtGGztfWJZW",
		"nvLzDTDFDFqSqncTFddLGfWfssVWWVdlGs",
		"njScSmcnFDDhnmcmnSBzMggjjMQrQCjBQrwH",
		"MFVFHqhHHfVVSGcVQCLttttWTtLQ",
		"BgdJJrJzbGpLssCtTLLCpC",
		"PGBGdjjllBGBjhmhHRlmSfhRDF",
		"BLJmWwBJDDmLDFnVPzvTttvNhNFsHhvvQH",
		"cjbSqfWbRrbSzNMjhtNzzTsT",
		"dRdbgflfqmWggGmwWW",
		"lBTTnDMnNwWdpwMllTMDdTBTSRJjRRcSJRhRGhwtccScgjtg",
		"vvsnCVnHHnsvPzLfVJjtShgGJGVJSc",
		"zQQvzCZbsnbHCWMdFpMqblBqND",
		"CLggQjStSQjLgVRfhBRztwpBpt",
		"DNVmJDFcGNGlmNDnGFnsGcDdwzzZwZPsBPZZhhfBTpwhPfZh",
		"DMDccrnrdnDJcJGFGmCqbLvQqHVSSQVMCgHV",
		"rZVNNDrCFCbjbRSgjhZv",
		"czcMTcGMlcjjvvvGdhbb",
		"pHMpHtLWHHHztMzsvqvnVqNvsnNDqW",
		"hPhPBQVjzjQScjChRVVVsdfbvdmvvpGmvfdbff",
		"nwZwZWNTrwLTTDtbfmHDpGccstDv",
		"FLFJNllFrwTgTwNLnwTlFPhcjjhRhSMSjjBPhzJPQq",
		"CgJCHgJfHzGrrMjJpl",
		"SWqQLSQqdFSLrrCSZrvpzjMM",
		"QLWqhFLQdwFqnPbHcCHPbhCh",
		"jLplfcfjQfptPtLTTtPrRqCCCjZvhBhwhDjwhBBBqC",
		"msznRWgSmmHbSJwDvvBqCCqWFF",
		"mzdsVnbbSSznmbGgGSmTclPcrQffdLfMTLRLtl",
		"PpQQFdNFBtdsFNNPPthTtldwGMnZVvmbbVGbMwrGvG",
		"wHWRJDjHHRgJDZlmVZmnvDbDDM",
		"JqJfgRWLSqWJfcsBfwPpPwFpNwfC",
		"SjpVgjghVZvssgsHjQfHcfcnfNcnbqcRbr",
		"WBtNWTWNJnCTCbRbQR",
		"FlWGtwPJmJPBmwFtMGmpNhDSSSgZhSszzghsMD",
		"fSfzvQSbbSSSTQQzDQTzHsqHmjHVFsjqVJsbrrLs",
		"GWZncGGdBwlPGPJcGlBcPgNGqLqmjsHMVLrrVMwjsMMssLmr",
		"ltGWNggZJnCRhDtzvtTt",
		"zzSHMgsPWzLSJQMMWQWLJBLVqcmVrzmvcpFcqmqpmFprFV",
		"TlDnwhDblbnbbtbjdpVCrmFVDgmprcRCcF",
		"hTdjgdldgdfZWWGGHQJHWZQH",
		"TZVsSRGsGMGWZTvpmrgcMCmrwwFFgF",
		"BDPPnDDLDqDFLNLgctLNrm",
		"PzDPfHPdDPdJPfdHJqdbnnSVvRvsjZSvbsrRGhvGRhTG",
		"nnghnhLhgdVqSPLHcPHPNtpmrRptNt",
		"DwvMWlsJlGzsGMlvsZcWWbRrNHNtZtHttHNTNttTrTTR",
		"DDzDcwlWWGlcsszGwBCggFqgghCBfqCFdhSh",
		"GzgQpJnwgJfbSWpSvqtvlBtmSLmLlvLS",
		"zCHsRZjHdzVPRFNlLlvtlNMNNrtDrt",
		"zzZVsCcCCHzRhcbpwGJGGpcf",
		"lVQMrwlMwwMCBZmFtthmVmsgWhTL",
		"bdzHSSFJcvzcpFDptsDDmhTgLmnmmmhT",
		"SFpcGvdpddvGlBPZMCGrBZ",
		"QmQTQTFTSQPNbsPjPnntZjjnnDlBBB",
		"JHqJqhfCJpWfpHchRzzCnGBjtjDZltsZpljGntrr",
		"chdMHqHcMhWMfRczMJmmQsVTSNbNNFbdQNST",
		"HHdFqqDRdNsHfNRsjWPCBcCCZPwDCZVBVc",
		"lhlgprMrlprbplzwZQPwwPjbZZCPwP",
		"lpljllTGGGglhThpjSvdssSnfRRTdRnvsN",
		"sDzjCqLLzddjsdRsSShgmtnCgFnmnffmmFcf",
		"rJbqJMTqJGVTrgnFmfhcfmnJgm",
		"GQbBWZGbVrqpqWpZZHlwDsDsRsdDlDRDBDsD",
		"LbLbvbhQgblwwqbjGG",
		"cFTTMsczJsTWJFfNNlVqvDqjNvlFqZ",
		"mfMvJHvmdLgLRgpHRR",
		"cvhRpWWhpfpcTpWvRcRcWVNwsjNLJFsJjwLtMLdsdsLJjL",
		"ZZPqGqgZrllbbVMtnJtllwJtltnj",
		"rrPCSQHbbrSqHqqZGGQbvzBVTpfzmvBvcvRCTpzT",
		"VMzNNhWVlrbJHbjcJCjCSR",
		"qgtDBgBtTGjqJvSPRJHCqHcd",
		"GfnfLjjgGLmgBWpWrMMnzWZppl",
		"zRtfBftCvvPSvPclZgcgJznWcgnq",
		"dDpGpVdwdGphbDMpdQhnJjjqnZQWnZNcNWlqlJ",
		"wdFdTDpGDSCmTtRqRq",
		"bTGrRzjbmbmqsGDDjHPpWpfjHJZFVPJp",
		"LwdwNNgMLfZTdCHPPd",
		"lhtwtvttgnchcsrvmzTvmsTbbm",
		"dfLjdlqLtqpbpPQpHS",
		"ZGJnFZFDpBWWGBTzzrhmhHzHQPFPvQ",
		"NgJGGGNGnJWgMDWGgDpWnZWgwdRCtwCCRVLdjwVcccdwct",
		"MMtzRCTnVVnHbbMNrHMRRVQJFrPPDsrsJgjjQGJGpFFF",
		"ZcqWqdfcmfhwDgGpDJmmQDQs",
		"LhwBddvlddvdLCgCtCTzgzMN",
		"ZffpWcfPcPrTwlVCvDhhcS",
		"GMLBNHjMBGtmjtMQMJjLHTwwsSswdslhGwDhsDvCvV",
		"tQgNRHMHQggHjQttbZqfFqrnqSZWfPZR",
		"MDqbPdqGfwGbfMswqfJdPGgQpCZvQgmvJFCmQJQvRCQv",
		"WFthczzzrpWgpHlRRC",
		"BhLSFnFrcLzhtSFnSTznhGbMGwjGwjjPbGDqjdqTMd",
		"zbQdJbBPTsWcCgdmfC",
		"DNqZHjvwZMShDhwqvLmlgVnLmcnsfnVf",
		"HHSFSjNZFqjZrhQpPtzrBgQzPrpB",
		"BzzJHvJJvWsgzPPTWPSJsJgNtmMtNFlvNZFltFtttjmtVZ",
		"cbccGnqhwhdpqRnnrdqdntZVFfMCltMtGFCmtFNFlj",
		"wnlRRQbnpHQWHgSQTs",
		"SQQBFnPzSnQVSzFSWlzlBSWFMsqmMmzLLChTmmMqzsTChmss",
		"ZgdwrJdtJrgpcCwZNbbjsGhGDvMrrqmmMmDhTsqv",
		"NtRZwCbZbbjRZNJcRbjWVWPnWHQWVPBQnHlfRB",
		"jLtlpljLpsbRnDNtBpbjdqWcqChccbhqChqbWQbQ",
		"gVwwggvJJwBqgWqfBCgT",
		"VHZrPSHwzSwvDLnnLlsrDRBp",
		"GBDGDrhwVrFhVCVBvhhvwBDVcmlfcPcMMmmmqNTScNNNflSF",
		"RzRLRbddjZTbTQjJSWNzlNWNcPqNWqSf",
		"dHjRgddHnZngjHLnQsgZHbGrrVpwpTtThnDBppVCTpDp",
		"dpjwvdLwtvJszhmzhRVj",
		"QfFQrffPBCBZMQrMRWzmzmVWVghVLs",
		"GBPGBrFrHZTBSTqvvLNtqSpq",
		"zFcGnQcZPZncbbbhPncpNCwzvwmjMvwwmwmpvd",
		"tBRtdrRTJNvRjMMwRC",
		"rlWSqVtTHqtVTrngQnfbQFdcghgW",
		"RMjfvsbQPvvNpLmLprFJFrFJJT",
		"SqCtGCcZZCwcCqqcGdWGdSZmMTmnVBdLnTTgTrMgrgVLTB",
		"lGzGzHSCzHctsQRMsQPzDjNs",
		"FQTWdTVMDWWVWFDTFcVcWJqTlCCCSlFmtCmNZStsmFmCwCgm",
		"nPLbGZzrgwBlbBlS",
		"PppGRnzPnpzvZVDJVvTfDVVVdD",
		"CQlDsWWfflWDMMhRmTGqFwSjJFdqwSQJzF",
		"PgpVbZPcHgBcgZrBZcHNTTbdqTGjSqFjRFjwSjGG",
		"VvgRBBNLRrPhsfhLWLWDtW",
		"hHhZDcDcPZWpLZWpCQ",
		"pbqdwmbqqmBpdMsgdqjCGvfCWGTWLGmWfGQtVt",
		"wjjbsFBRbwFgMpDznPDrrFcShr",
		"zGmsJQsDBBmDDJJpRZzSqZZPRffWRSqh",
		"LQlVHjCCNCLRSgWlZnhPZW",
		"bLtHCHtjVHQDTJcGDQbs",
		"QPRlnHfPllgRfnqhcwgwGMGzBWzBBBBB",
		"LCtVCZtDbdttjZFtvjvdDMGGBmBWWZWBBWGSsWJSJm",
		"dFVNTVbTjMnHlThHnpHR",
		"hGhZLlqmqZWTcWrmWqrWmTrqjQVQwNHPgPwPPSgPjNwVSLjD",
		"bsMJBRMvRsvsJMRRbspFgQwSNNSNwVNjgwDMDgHH",
		"pFnFtvCQFsbBQzsBrmnTWchWZqWqWqGZ",
		"wpwlJdCWJWdzlWGRRcrDrwRqrqDFrF",
		"ZmPsSVnnTvmHDgFcDTHFTF",
		"smVvnSQhbsNsvsmsnQQclzGCWpphppLJCdJWBpll",
		"FfSmMJmBDfBjDSjFtBVmftBQPwPhPCbhQvbhwCCbvhhjhq",
		"ZnZHZgclJWNbwbcbcwPrQv",
		"ZNNGLHzHNZTTJnlFFfVmDBsFSVLdSD",
		"DDBvjMvBJMtWjNRNrvdtbshTdpssdPSgpFpP",
		"LwGQcwLGJTSssQbg",
		"wcfcCcmHfJmcLLVZVqBBBqWDRqBRDWzHNN",
		"LhnpcdcdrChplvllHptlgR",
		"TsSTSBqPBTGNzqGfzfffGfVtPlHHvMDtHtRMlvjHHgtv",
		"sBfWNGQmQbCgQhQn",
		"pnnHngqsRjstjRgtrBDlBwDgwDZBldlD",
		"SFvJMGhhvcbMSWPJbFzVDdzDZwDFlFlzfFBV",
		"JJNWGSWSMNMCPhcvMhGStwtjtRRtRRqRsCtQjqRj",
		"nQZfJswFffjJplqhlqZlVVhp",
		"vtdCvGdBCzVHgnzLDlpL",
		"BbCTGvtGnBBCPjrsccjrwbrFjj",
		"wjjvDQwWvSQDLbwfNrrJcMHrczcpWN",
		"tTnqlRsTRthFhnnpfNrmcTNMzfCzHr",
		"slsBGVlqghMRhsRlltsFDgQSDbQwPSPLQSZjPSLQ",
		"PcQmmQRQZRFQPjjDWhgCgWdM",
		"nBGtbGNBBGvJbbtpWhCNCjmMjHlWHDdM",
		"vbBnBBrrvVBtJJbvtzzptRQLLZLRFfqLSTSfzqwfqm",
		"PBFZDFPsDZBnTTBdDHSNSpNzmVTVmVGlNH",
		"WvqFLWQCMQLMRtQJFWQLvQCJVzmNjNzVCllVmGSlNHlVzljH",
		"WRtrhtWQWFbrrBwBZrbn",
		"MwWnMPnMPNswjPDvRbsblCGFZGZF",
		"JdJVVVtLdgZhvGVBDZhZ",
		"rqmdqtgcLQLTfWffwGGmmp",
		"QQhhWzQWsMhZjbbmffgfjrGDdfdGvv",
		"HnCJVHcnnHttCRVRCcnBCqJBGfPmTvTvdDfvqfrddNDDggGG",
		"wBwwcRBBCJpFcFcpnVVFQWQLSLbQZLmhzshMlMwQ",
		"CgDNbsDcHjTcgDCgjRHMJrlHFrBHFmrttrGGtFwG",
		"VfQJnvJdhvSJZphSzWpvSzZSltGGBllGBqPPFPrwrmfqtFFB",
		"VJLvdQhphhQnjLsjDDjDcRbL",
		"LjljTPdLdccLhcMZhTTMdzrrtzGgtvrgnttNDGrWtDgn",
		"hbCmCHqHmSbRgNrtvCgGgttN",
		"SJFJRFpHqwSFSpsHwbHwRhSJPTjMMMPdPlPLcVQscLVQQVlL",
		"QmTTQVqrVrMvbCwLczbRlQ",
		"sSNtNGZFjBSsjSSShcRvwCFLlzWcWRzCWv",
		"LLNjZhSGZBnjhJjjrrTgPqgPgTfqfJpf",
		"DCCjDDtHVptCtvMZlJbSnScWWfHlhW",
		"qsTFmTgmqRRLswQGmfWwSnZSSfJSSWZcWb",
		"dsmqgqdsNgTFLFRLGmRpBtBDtDNVpPCCfVrtpD",
		"LLNRhHhRbsNGjqCBPBrLCw",
		"lgcFfvWGTllzfJVVFVWDzFqqMrZCMBrZZqvSrCPZSSrr",
		"fFGGlTTTlzWQGzzDFttQmHnpnhtmmpRssm",
		"LZNnBgtlNZztzmGHmpHHPPPm",
		"QwjjQRCQScbFFFchhFrFjwmsNPHSSWJGsGppMSWMmWqs",
		"dQCwwbwhrjQQjCwRwbhRBlDDfBtVlnNnlgLdnvvd",
		"wRbGbqqGCwnGTRqBqlMVphpgMgMFdnFVnt",
		"cgzvssscgHJVdhDdhDMvDM",
		"PjcZcsJrJHWgrPBQmCqRBPSSRCSb",
		"rHBmLlPLlTzztvRtGsVL",
		"NWJJWWjJDJWWhphqwCFCwzvRVcgRtctRgNNVVscsGc",
		"hqCCsnCpCDnbCnWhwpbHbBHmMBMMmdPrZfdP",
		"GRPprPQdsprGpGgGTlqfVqnZLgnLnwNZLw",
		"CCWJMMvhhCvthtCjvDWFjMcCVZJLNnfqnllwzlzNnzzwVNnN",
		"cDtZFjDjcMCDDtZFSMCvvDpmsmSRRpmmbSpdPRdTmTsp",
		"mmPpsbZZbbzvzgbrZRPgPMWqtHtqDnGCnMWCDwGHwtwW",
		"cBFBNNccsTLjJjfcjfGDGQtWwFCnCGtqCCQH",
		"TNsTLJlffdldzvrmbmrPzp"}

	return array
}
