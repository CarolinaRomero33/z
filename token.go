package andrea

// Token represents a lexical token.
type Token int

const (
	
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT // main
	// Misc characters
	ASTERISK      // *
	COMMA         // ,
	DosPunt       // :
	ParDer        // (
	ParIsq        // )
	PuntCom       // ;
	LlaveDer      // {
	LlaveIsq      // }
	SaltoDeL      //n
	Asignaicon    // =
	corchetederecho // [
	corcheteizquierdo // ]

	// Keywords
	zselect
	zcos
	ztan
	zsen


)
