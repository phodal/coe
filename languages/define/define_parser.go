// Code generated from Define.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Define

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 126,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 7, 2, 34, 10, 2,
	12, 2, 14, 2, 37, 11, 2, 3, 2, 5, 2, 40, 10, 2, 3, 2, 7, 2, 43, 10, 2,
	12, 2, 14, 2, 46, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	5, 4, 56, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 62, 10, 5, 12, 5, 14, 5,
	65, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 82, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 7, 9, 91, 10, 9, 12, 9, 14, 9, 94, 11, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 124, 10, 16, 3, 16, 2, 2, 17,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 2, 2, 118, 2,
	35, 3, 2, 2, 2, 4, 49, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2, 8, 57, 3, 2, 2, 2,
	10, 81, 3, 2, 2, 2, 12, 83, 3, 2, 2, 2, 14, 85, 3, 2, 2, 2, 16, 87, 3,
	2, 2, 2, 18, 98, 3, 2, 2, 2, 20, 100, 3, 2, 2, 2, 22, 104, 3, 2, 2, 2,
	24, 106, 3, 2, 2, 2, 26, 108, 3, 2, 2, 2, 28, 114, 3, 2, 2, 2, 30, 123,
	3, 2, 2, 2, 32, 34, 5, 4, 3, 2, 33, 32, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2,
	35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3,
	2, 2, 2, 38, 40, 5, 8, 5, 2, 39, 38, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40,
	44, 3, 2, 2, 2, 41, 43, 5, 6, 4, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2,
	2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44,
	3, 2, 2, 2, 47, 48, 7, 2, 2, 3, 48, 3, 3, 2, 2, 2, 49, 50, 7, 3, 2, 2,
	50, 51, 7, 11, 2, 2, 51, 52, 7, 25, 2, 2, 52, 5, 3, 2, 2, 2, 53, 56, 5,
	20, 11, 2, 54, 56, 5, 26, 14, 2, 55, 53, 3, 2, 2, 2, 55, 54, 3, 2, 2, 2,
	56, 7, 3, 2, 2, 2, 57, 58, 7, 5, 2, 2, 58, 59, 5, 22, 12, 2, 59, 63, 7,
	14, 2, 2, 60, 62, 5, 10, 6, 2, 61, 60, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2,
	63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 63, 3,
	2, 2, 2, 66, 67, 7, 15, 2, 2, 67, 9, 3, 2, 2, 2, 68, 69, 7, 6, 2, 2, 69,
	70, 7, 21, 2, 2, 70, 71, 5, 12, 7, 2, 71, 72, 5, 14, 8, 2, 72, 82, 3, 2,
	2, 2, 73, 74, 7, 7, 2, 2, 74, 75, 7, 12, 2, 2, 75, 76, 7, 11, 2, 2, 76,
	77, 7, 13, 2, 2, 77, 78, 7, 14, 2, 2, 78, 79, 5, 16, 9, 2, 79, 80, 7, 15,
	2, 2, 80, 82, 3, 2, 2, 2, 81, 68, 3, 2, 2, 2, 81, 73, 3, 2, 2, 2, 82, 11,
	3, 2, 2, 2, 83, 84, 7, 11, 2, 2, 84, 13, 3, 2, 2, 2, 85, 86, 7, 11, 2,
	2, 86, 15, 3, 2, 2, 2, 87, 88, 5, 12, 7, 2, 88, 92, 7, 11, 2, 2, 89, 91,
	5, 12, 7, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2,
	92, 93, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 96, 5,
	18, 10, 2, 96, 97, 5, 12, 7, 2, 97, 17, 3, 2, 2, 2, 98, 99, 7, 25, 2, 2,
	99, 19, 3, 2, 2, 2, 100, 101, 5, 22, 12, 2, 101, 102, 7, 21, 2, 2, 102,
	103, 5, 24, 13, 2, 103, 21, 3, 2, 2, 2, 104, 105, 7, 11, 2, 2, 105, 23,
	3, 2, 2, 2, 106, 107, 7, 11, 2, 2, 107, 25, 3, 2, 2, 2, 108, 109, 7, 8,
	2, 2, 109, 110, 7, 11, 2, 2, 110, 111, 7, 14, 2, 2, 111, 112, 5, 28, 15,
	2, 112, 113, 7, 15, 2, 2, 113, 27, 3, 2, 2, 2, 114, 115, 7, 11, 2, 2, 115,
	116, 7, 14, 2, 2, 116, 117, 5, 30, 16, 2, 117, 118, 7, 15, 2, 2, 118, 29,
	3, 2, 2, 2, 119, 120, 7, 9, 2, 2, 120, 124, 7, 25, 2, 2, 121, 122, 7, 10,
	2, 2, 122, 124, 7, 25, 2, 2, 123, 119, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2,
	124, 31, 3, 2, 2, 2, 10, 35, 39, 44, 55, 63, 81, 92, 123,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'symbol'", "", "'define'", "'defaultSymbol'", "'defaultTemplate'",
	"'module'", "'import'", "'equal'", "", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "';'", "','", "'.'", "':'",
}
var symbolicNames = []string{
	"", "SYMBOL_TEXT", "SPECIAL_SYMBOL", "DEFINE", "DEFAULT_SYMBOL", "DEFAULT_TEMPLATE",
	"MODULE", "IMPORT", "EQUAL", "IDENTIFIER", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "COLON", "WS", "COMMENT",
	"LINE_COMMENT", "STRING_LITERAL", "Symbols",
}

var ruleNames = []string{
	"compilationUnit", "symbolDelcaration", "normalDeclarations", "defineDeclaration",
	"defineExpress", "symbolKey", "symbolValue", "defineBody", "templateData",
	"systemDeclaration", "defineKey", "defineValue", "moduleDeclaration", "moduleDefine",
	"moduleAttributes",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DefineParser struct {
	*antlr.BaseParser
}

func NewDefineParser(input antlr.TokenStream) *DefineParser {
	this := new(DefineParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Define.g4"

	return this
}

// DefineParser tokens.
const (
	DefineParserEOF              = antlr.TokenEOF
	DefineParserSYMBOL_TEXT      = 1
	DefineParserSPECIAL_SYMBOL   = 2
	DefineParserDEFINE           = 3
	DefineParserDEFAULT_SYMBOL   = 4
	DefineParserDEFAULT_TEMPLATE = 5
	DefineParserMODULE           = 6
	DefineParserIMPORT           = 7
	DefineParserEQUAL            = 8
	DefineParserIDENTIFIER       = 9
	DefineParserLPAREN           = 10
	DefineParserRPAREN           = 11
	DefineParserLBRACE           = 12
	DefineParserRBRACE           = 13
	DefineParserLBRACK           = 14
	DefineParserRBRACK           = 15
	DefineParserSEMI             = 16
	DefineParserCOMMA            = 17
	DefineParserDOT              = 18
	DefineParserCOLON            = 19
	DefineParserWS               = 20
	DefineParserCOMMENT          = 21
	DefineParserLINE_COMMENT     = 22
	DefineParserSTRING_LITERAL   = 23
	DefineParserSymbols          = 24
)

// DefineParser rules.
const (
	DefineParserRULE_compilationUnit    = 0
	DefineParserRULE_symbolDelcaration  = 1
	DefineParserRULE_normalDeclarations = 2
	DefineParserRULE_defineDeclaration  = 3
	DefineParserRULE_defineExpress      = 4
	DefineParserRULE_symbolKey          = 5
	DefineParserRULE_symbolValue        = 6
	DefineParserRULE_defineBody         = 7
	DefineParserRULE_templateData       = 8
	DefineParserRULE_systemDeclaration  = 9
	DefineParserRULE_defineKey          = 10
	DefineParserRULE_defineValue        = 11
	DefineParserRULE_moduleDeclaration  = 12
	DefineParserRULE_moduleDefine       = 13
	DefineParserRULE_moduleAttributes   = 14
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(DefineParserEOF, 0)
}

func (s *CompilationUnitContext) AllSymbolDelcaration() []ISymbolDelcarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolDelcarationContext)(nil)).Elem())
	var tst = make([]ISymbolDelcarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolDelcarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) SymbolDelcaration(i int) ISymbolDelcarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolDelcarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolDelcarationContext)
}

func (s *CompilationUnitContext) DefineDeclaration() IDefineDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineDeclarationContext)
}

func (s *CompilationUnitContext) AllNormalDeclarations() []INormalDeclarationsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INormalDeclarationsContext)(nil)).Elem())
	var tst = make([]INormalDeclarationsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INormalDeclarationsContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) NormalDeclarations(i int) INormalDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormalDeclarationsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INormalDeclarationsContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DefineParserRULE_compilationUnit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserSYMBOL_TEXT {
		{
			p.SetState(30)
			p.SymbolDelcaration()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DefineParserDEFINE {
		{
			p.SetState(36)
			p.DefineDeclaration()
		}

	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserMODULE || _la == DefineParserIDENTIFIER {
		{
			p.SetState(39)
			p.NormalDeclarations()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
		p.Match(DefineParserEOF)
	}

	return localctx
}

// ISymbolDelcarationContext is an interface to support dynamic dispatch.
type ISymbolDelcarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolDelcarationContext differentiates from other interfaces.
	IsSymbolDelcarationContext()
}

type SymbolDelcarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolDelcarationContext() *SymbolDelcarationContext {
	var p = new(SymbolDelcarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_symbolDelcaration
	return p
}

func (*SymbolDelcarationContext) IsSymbolDelcarationContext() {}

func NewSymbolDelcarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolDelcarationContext {
	var p = new(SymbolDelcarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_symbolDelcaration

	return p
}

func (s *SymbolDelcarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolDelcarationContext) SYMBOL_TEXT() antlr.TerminalNode {
	return s.GetToken(DefineParserSYMBOL_TEXT, 0)
}

func (s *SymbolDelcarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *SymbolDelcarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *SymbolDelcarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolDelcarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolDelcarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSymbolDelcaration(s)
	}
}

func (s *SymbolDelcarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSymbolDelcaration(s)
	}
}

func (s *SymbolDelcarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSymbolDelcaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SymbolDelcaration() (localctx ISymbolDelcarationContext) {
	localctx = NewSymbolDelcarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DefineParserRULE_symbolDelcaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(DefineParserSYMBOL_TEXT)
	}
	{
		p.SetState(48)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(49)
		p.Match(DefineParserSTRING_LITERAL)
	}

	return localctx
}

// INormalDeclarationsContext is an interface to support dynamic dispatch.
type INormalDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNormalDeclarationsContext differentiates from other interfaces.
	IsNormalDeclarationsContext()
}

type NormalDeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormalDeclarationsContext() *NormalDeclarationsContext {
	var p = new(NormalDeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_normalDeclarations
	return p
}

func (*NormalDeclarationsContext) IsNormalDeclarationsContext() {}

func NewNormalDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalDeclarationsContext {
	var p = new(NormalDeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_normalDeclarations

	return p
}

func (s *NormalDeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalDeclarationsContext) SystemDeclaration() ISystemDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISystemDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISystemDeclarationContext)
}

func (s *NormalDeclarationsContext) ModuleDeclaration() IModuleDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleDeclarationContext)
}

func (s *NormalDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalDeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterNormalDeclarations(s)
	}
}

func (s *NormalDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitNormalDeclarations(s)
	}
}

func (s *NormalDeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitNormalDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) NormalDeclarations() (localctx INormalDeclarationsContext) {
	localctx = NewNormalDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DefineParserRULE_normalDeclarations)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefineParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.SystemDeclaration()
		}

	case DefineParserMODULE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.ModuleDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefineDeclarationContext is an interface to support dynamic dispatch.
type IDefineDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineDeclarationContext differentiates from other interfaces.
	IsDefineDeclarationContext()
}

type DefineDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineDeclarationContext() *DefineDeclarationContext {
	var p = new(DefineDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineDeclaration
	return p
}

func (*DefineDeclarationContext) IsDefineDeclarationContext() {}

func NewDefineDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineDeclarationContext {
	var p = new(DefineDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineDeclaration

	return p
}

func (s *DefineDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineDeclarationContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(DefineParserDEFINE, 0)
}

func (s *DefineDeclarationContext) DefineKey() IDefineKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineKeyContext)
}

func (s *DefineDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *DefineDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *DefineDeclarationContext) AllDefineExpress() []IDefineExpressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefineExpressContext)(nil)).Elem())
	var tst = make([]IDefineExpressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefineExpressContext)
		}
	}

	return tst
}

func (s *DefineDeclarationContext) DefineExpress(i int) IDefineExpressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineExpressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefineExpressContext)
}

func (s *DefineDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineDeclaration(s)
	}
}

func (s *DefineDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineDeclaration(s)
	}
}

func (s *DefineDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineDeclaration() (localctx IDefineDeclarationContext) {
	localctx = NewDefineDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DefineParserRULE_defineDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(DefineParserDEFINE)
	}
	{
		p.SetState(56)
		p.DefineKey()
	}
	{
		p.SetState(57)
		p.Match(DefineParserLBRACE)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserDEFAULT_SYMBOL || _la == DefineParserDEFAULT_TEMPLATE {
		{
			p.SetState(58)
			p.DefineExpress()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IDefineExpressContext is an interface to support dynamic dispatch.
type IDefineExpressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineExpressContext differentiates from other interfaces.
	IsDefineExpressContext()
}

type DefineExpressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineExpressContext() *DefineExpressContext {
	var p = new(DefineExpressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineExpress
	return p
}

func (*DefineExpressContext) IsDefineExpressContext() {}

func NewDefineExpressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineExpressContext {
	var p = new(DefineExpressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineExpress

	return p
}

func (s *DefineExpressContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineExpressContext) DEFAULT_SYMBOL() antlr.TerminalNode {
	return s.GetToken(DefineParserDEFAULT_SYMBOL, 0)
}

func (s *DefineExpressContext) COLON() antlr.TerminalNode {
	return s.GetToken(DefineParserCOLON, 0)
}

func (s *DefineExpressContext) SymbolKey() ISymbolKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolKeyContext)
}

func (s *DefineExpressContext) SymbolValue() ISymbolValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolValueContext)
}

func (s *DefineExpressContext) DEFAULT_TEMPLATE() antlr.TerminalNode {
	return s.GetToken(DefineParserDEFAULT_TEMPLATE, 0)
}

func (s *DefineExpressContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DefineParserLPAREN, 0)
}

func (s *DefineExpressContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *DefineExpressContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DefineParserRPAREN, 0)
}

func (s *DefineExpressContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *DefineExpressContext) DefineBody() IDefineBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineBodyContext)
}

func (s *DefineExpressContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *DefineExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineExpressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineExpressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineExpress(s)
	}
}

func (s *DefineExpressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineExpress(s)
	}
}

func (s *DefineExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineExpress() (localctx IDefineExpressContext) {
	localctx = NewDefineExpressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DefineParserRULE_defineExpress)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefineParserDEFAULT_SYMBOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(DefineParserDEFAULT_SYMBOL)
		}
		{
			p.SetState(67)
			p.Match(DefineParserCOLON)
		}
		{
			p.SetState(68)
			p.SymbolKey()
		}
		{
			p.SetState(69)
			p.SymbolValue()
		}

	case DefineParserDEFAULT_TEMPLATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(DefineParserDEFAULT_TEMPLATE)
		}
		{
			p.SetState(72)
			p.Match(DefineParserLPAREN)
		}
		{
			p.SetState(73)
			p.Match(DefineParserIDENTIFIER)
		}
		{
			p.SetState(74)
			p.Match(DefineParserRPAREN)
		}
		{
			p.SetState(75)
			p.Match(DefineParserLBRACE)
		}
		{
			p.SetState(76)
			p.DefineBody()
		}
		{
			p.SetState(77)
			p.Match(DefineParserRBRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISymbolKeyContext is an interface to support dynamic dispatch.
type ISymbolKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolKeyContext differentiates from other interfaces.
	IsSymbolKeyContext()
}

type SymbolKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolKeyContext() *SymbolKeyContext {
	var p = new(SymbolKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_symbolKey
	return p
}

func (*SymbolKeyContext) IsSymbolKeyContext() {}

func NewSymbolKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolKeyContext {
	var p = new(SymbolKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_symbolKey

	return p
}

func (s *SymbolKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolKeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *SymbolKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSymbolKey(s)
	}
}

func (s *SymbolKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSymbolKey(s)
	}
}

func (s *SymbolKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSymbolKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SymbolKey() (localctx ISymbolKeyContext) {
	localctx = NewSymbolKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DefineParserRULE_symbolKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// ISymbolValueContext is an interface to support dynamic dispatch.
type ISymbolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolValueContext differentiates from other interfaces.
	IsSymbolValueContext()
}

type SymbolValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolValueContext() *SymbolValueContext {
	var p = new(SymbolValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_symbolValue
	return p
}

func (*SymbolValueContext) IsSymbolValueContext() {}

func NewSymbolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolValueContext {
	var p = new(SymbolValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_symbolValue

	return p
}

func (s *SymbolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *SymbolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSymbolValue(s)
	}
}

func (s *SymbolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSymbolValue(s)
	}
}

func (s *SymbolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSymbolValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SymbolValue() (localctx ISymbolValueContext) {
	localctx = NewSymbolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DefineParserRULE_symbolValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// IDefineBodyContext is an interface to support dynamic dispatch.
type IDefineBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineBodyContext differentiates from other interfaces.
	IsDefineBodyContext()
}

type DefineBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineBodyContext() *DefineBodyContext {
	var p = new(DefineBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineBody
	return p
}

func (*DefineBodyContext) IsDefineBodyContext() {}

func NewDefineBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineBodyContext {
	var p = new(DefineBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineBody

	return p
}

func (s *DefineBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineBodyContext) AllSymbolKey() []ISymbolKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolKeyContext)(nil)).Elem())
	var tst = make([]ISymbolKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolKeyContext)
		}
	}

	return tst
}

func (s *DefineBodyContext) SymbolKey(i int) ISymbolKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolKeyContext)
}

func (s *DefineBodyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *DefineBodyContext) TemplateData() ITemplateDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateDataContext)
}

func (s *DefineBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineBody(s)
	}
}

func (s *DefineBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineBody(s)
	}
}

func (s *DefineBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineBody() (localctx IDefineBodyContext) {
	localctx = NewDefineBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DefineParserRULE_defineBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.SymbolKey()
	}
	{
		p.SetState(86)
		p.Match(DefineParserIDENTIFIER)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DefineParserIDENTIFIER {
		{
			p.SetState(87)
			p.SymbolKey()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
		p.TemplateData()
	}
	{
		p.SetState(94)
		p.SymbolKey()
	}

	return localctx
}

// ITemplateDataContext is an interface to support dynamic dispatch.
type ITemplateDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateDataContext differentiates from other interfaces.
	IsTemplateDataContext()
}

type TemplateDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateDataContext() *TemplateDataContext {
	var p = new(TemplateDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_templateData
	return p
}

func (*TemplateDataContext) IsTemplateDataContext() {}

func NewTemplateDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateDataContext {
	var p = new(TemplateDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_templateData

	return p
}

func (s *TemplateDataContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateDataContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *TemplateDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterTemplateData(s)
	}
}

func (s *TemplateDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitTemplateData(s)
	}
}

func (s *TemplateDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitTemplateData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) TemplateData() (localctx ITemplateDataContext) {
	localctx = NewTemplateDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DefineParserRULE_templateData)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(DefineParserSTRING_LITERAL)
	}

	return localctx
}

// ISystemDeclarationContext is an interface to support dynamic dispatch.
type ISystemDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSystemDeclarationContext differentiates from other interfaces.
	IsSystemDeclarationContext()
}

type SystemDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystemDeclarationContext() *SystemDeclarationContext {
	var p = new(SystemDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_systemDeclaration
	return p
}

func (*SystemDeclarationContext) IsSystemDeclarationContext() {}

func NewSystemDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SystemDeclarationContext {
	var p = new(SystemDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_systemDeclaration

	return p
}

func (s *SystemDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SystemDeclarationContext) DefineKey() IDefineKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineKeyContext)
}

func (s *SystemDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(DefineParserCOLON, 0)
}

func (s *SystemDeclarationContext) DefineValue() IDefineValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineValueContext)
}

func (s *SystemDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystemDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystemDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterSystemDeclaration(s)
	}
}

func (s *SystemDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitSystemDeclaration(s)
	}
}

func (s *SystemDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitSystemDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) SystemDeclaration() (localctx ISystemDeclarationContext) {
	localctx = NewSystemDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DefineParserRULE_systemDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.DefineKey()
	}
	{
		p.SetState(99)
		p.Match(DefineParserCOLON)
	}
	{
		p.SetState(100)
		p.DefineValue()
	}

	return localctx
}

// IDefineKeyContext is an interface to support dynamic dispatch.
type IDefineKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineKeyContext differentiates from other interfaces.
	IsDefineKeyContext()
}

type DefineKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineKeyContext() *DefineKeyContext {
	var p = new(DefineKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineKey
	return p
}

func (*DefineKeyContext) IsDefineKeyContext() {}

func NewDefineKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineKeyContext {
	var p = new(DefineKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineKey

	return p
}

func (s *DefineKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineKeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *DefineKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineKey(s)
	}
}

func (s *DefineKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineKey(s)
	}
}

func (s *DefineKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineKey() (localctx IDefineKeyContext) {
	localctx = NewDefineKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DefineParserRULE_defineKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// IDefineValueContext is an interface to support dynamic dispatch.
type IDefineValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineValueContext differentiates from other interfaces.
	IsDefineValueContext()
}

type DefineValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineValueContext() *DefineValueContext {
	var p = new(DefineValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_defineValue
	return p
}

func (*DefineValueContext) IsDefineValueContext() {}

func NewDefineValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineValueContext {
	var p = new(DefineValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_defineValue

	return p
}

func (s *DefineValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *DefineValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterDefineValue(s)
	}
}

func (s *DefineValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitDefineValue(s)
	}
}

func (s *DefineValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitDefineValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) DefineValue() (localctx IDefineValueContext) {
	localctx = NewDefineValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DefineParserRULE_defineValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(DefineParserIDENTIFIER)
	}

	return localctx
}

// IModuleDeclarationContext is an interface to support dynamic dispatch.
type IModuleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleDeclarationContext differentiates from other interfaces.
	IsModuleDeclarationContext()
}

type ModuleDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDeclarationContext() *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleDeclaration
	return p
}

func (*ModuleDeclarationContext) IsModuleDeclarationContext() {}

func NewModuleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleDeclaration

	return p
}

func (s *ModuleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDeclarationContext) MODULE() antlr.TerminalNode {
	return s.GetToken(DefineParserMODULE, 0)
}

func (s *ModuleDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *ModuleDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *ModuleDeclarationContext) ModuleDefine() IModuleDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleDefineContext)
}

func (s *ModuleDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *ModuleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleDeclaration() (localctx IModuleDeclarationContext) {
	localctx = NewModuleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DefineParserRULE_moduleDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(DefineParserMODULE)
	}
	{
		p.SetState(107)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(108)
		p.Match(DefineParserLBRACE)
	}
	{
		p.SetState(109)
		p.ModuleDefine()
	}
	{
		p.SetState(110)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IModuleDefineContext is an interface to support dynamic dispatch.
type IModuleDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleDefineContext differentiates from other interfaces.
	IsModuleDefineContext()
}

type ModuleDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDefineContext() *ModuleDefineContext {
	var p = new(ModuleDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleDefine
	return p
}

func (*ModuleDefineContext) IsModuleDefineContext() {}

func NewModuleDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDefineContext {
	var p = new(ModuleDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleDefine

	return p
}

func (s *ModuleDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDefineContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DefineParserIDENTIFIER, 0)
}

func (s *ModuleDefineContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserLBRACE, 0)
}

func (s *ModuleDefineContext) ModuleAttributes() IModuleAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleAttributesContext)
}

func (s *ModuleDefineContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DefineParserRBRACE, 0)
}

func (s *ModuleDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleDefine(s)
	}
}

func (s *ModuleDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleDefine(s)
	}
}

func (s *ModuleDefineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleDefine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleDefine() (localctx IModuleDefineContext) {
	localctx = NewModuleDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DefineParserRULE_moduleDefine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(DefineParserIDENTIFIER)
	}
	{
		p.SetState(113)
		p.Match(DefineParserLBRACE)
	}
	{
		p.SetState(114)
		p.ModuleAttributes()
	}
	{
		p.SetState(115)
		p.Match(DefineParserRBRACE)
	}

	return localctx
}

// IModuleAttributesContext is an interface to support dynamic dispatch.
type IModuleAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleAttributesContext differentiates from other interfaces.
	IsModuleAttributesContext()
}

type ModuleAttributesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleAttributesContext() *ModuleAttributesContext {
	var p = new(ModuleAttributesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DefineParserRULE_moduleAttributes
	return p
}

func (*ModuleAttributesContext) IsModuleAttributesContext() {}

func NewModuleAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleAttributesContext {
	var p = new(ModuleAttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DefineParserRULE_moduleAttributes

	return p
}

func (s *ModuleAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleAttributesContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(DefineParserIMPORT, 0)
}

func (s *ModuleAttributesContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DefineParserSTRING_LITERAL, 0)
}

func (s *ModuleAttributesContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(DefineParserEQUAL, 0)
}

func (s *ModuleAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.EnterModuleAttributes(s)
	}
}

func (s *ModuleAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DefineListener); ok {
		listenerT.ExitModuleAttributes(s)
	}
}

func (s *ModuleAttributesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DefineVisitor:
		return t.VisitModuleAttributes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DefineParser) ModuleAttributes() (localctx IModuleAttributesContext) {
	localctx = NewModuleAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DefineParserRULE_moduleAttributes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DefineParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(DefineParserIMPORT)
		}
		{
			p.SetState(118)
			p.Match(DefineParserSTRING_LITERAL)
		}

	case DefineParserEQUAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Match(DefineParserEQUAL)
		}
		{
			p.SetState(120)
			p.Match(DefineParserSTRING_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
