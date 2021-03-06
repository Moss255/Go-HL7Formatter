package main

func TestLoadFile(t *testing.T) {
	LoadFromFile
}

func TestFormatMessage(t *testing.T) {

	var segments []string = []string{"PID", "PD1", "PV1", "NTE", "PV2", "ZVI", "AL1", "ZAL", "ORC", "RXO", "RXR", "RXE", "OBX"}

	var text string = "MSH|^~\&|MILLENNIUM|WSH|WSH_TIE|WSH_TIE|20200310095702||RDE^O01|Q105132700T129367833||2.3||||||8859/1PID|1|1516070^^^RGR50^MRN|1516070^^^RGR50^MR~^^^NHS^NH||ZZZTEST^MEDS TEAM OAK^^^PROFESSOR^^CURRENT||19800101|2|||No Fixed Abode^^^""^ZZ99 3VZ^GBR^HOME^^""||||""|""|""|10012589^^^Encounter Num^FINNBR||||C||""|0|""|""|""||""PD1|""|""|Practice Code not known^^8023|G1234567^Test^GP^^^^^^External Id^PRSNL^^^EXTID^""|""||""|""PV1|1|INPATIENT|ZZZTEST LAB^Bay 1^Bed 10^WSH^^BED^Main Building|22|||C4598938^Suresh^Mohanraj^^^Dr^^^NHS Consultant Number^PRSNL^^^NONGP^""~716337131016^Suresh^Mohanraj^^^Dr^^^Doctor Nbr^PRSNL^^^ORGDR^""|G1234567^Test^GP^^^^^^External Id^PRSNL^^^EXTID^""||300|""|""|""|19|""|""||INPATIENT|10012589^^^Attendance Num^VISITID|""||""||||||||||||||""|""|""|WSH||ACTIVE|||20190308090900PV2||1||""|||""|||0|||""||||||||""|""|^^38024||||""ZVI|""|||||430|||""|""||""||""|20200107144700|""|""|""|""|""|""|""||""|""|""AL1|1|DRUG|##NOMEN##,AL1,ceStruct,allergy,32239,6334447^NKA^ALLERGYZAL|SNAPSHOT|20190308092025|1294449|1294449|ALLEGY|CANC|||||20190308092025|^Bognar^Julia^^^^^^^PRSNL|0ORC|CA|31320845^HNAM_ORDERID|||CA||^twice a day^^20200310180000^20200310095701^ROUTINE||20200310095641|^Reed^Matt^^^^^^^PRSNL||C4598938^Suresh^Mohanraj^^^Dr^^^NHS Consultant Number^PRSNL^^^NONGP^""~716337131016^Suresh^Mohanraj^^^Dr^^^Doctor Nbr^PRSNL^^^ORGDR^""|""||20200310095701|Course completed^Course completedRXO|CD:14490226^Diltiazem^NHS_LEG_INT^^Diltiazem|120||mg^^^^120 mg / 1 capsule|CD:14093882^capsule (modified release)||||||0|""|0RXR|PO^oral|||CD:10915RXE|^twice a day^""^20200310180000^20200310095701^ROUTINE|##ITEM##,IDENTIFIER_TYPE_CD,VALUE_KEY,524232,1,CE,coding_system^##ITEM##,DESC,VALUE,524232,1^^^Diltiazem|120||mg^^^^120 mg / 1 capsule|CD:14093882^capsule (modified release)||||1|EA|||^Reed^Matt^^^^^^^PRSNL|31320845RXR|PO^oral|||CD:10915OBX|1|TS|REQUESTED START DATE/TIME^Requested Start Date/Time^NHS_LEG_INT||20200310180000OBX|2|ST|IV SET SHELL ITEM ID^IV Set Shell Item Id^NHS_LEG_INT||0OBX|3|TS|CANCEL DATE AND TIME^Cancel Date/Time^NHS_LEG_INT||20200310095600"

	text := FormatMessage(text, segments)

	
	
}

func TestAddNewLine(t *testing.T) {

}

func TestWriteFile(t *testing.T) {

}
