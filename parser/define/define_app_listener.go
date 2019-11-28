package code

import (
	. "../../languages/define"
	"reflect"
	"strings"
)

type Symbol struct {
	Key   string
	Value string
}

var symbols []Symbol
var symbolMaps map[string]string

var startSymbol string
var endSymbol string

type DefineInformation struct {
	DefineTemplates map[string]string
}

var defineInformation DefineInformation

func NewDefineAppListener(start string, end string) *DefineAppListener {
	startSymbol = start
	endSymbol = end
	defineInformation = *&DefineInformation{make(map[string]string)}
	symbolMaps = make(map[string]string)
	return &DefineAppListener{}
}

type DefineAppListener struct {
	BaseDefineListener
}

func (s *DefineAppListener) EnterSymbolDeclaration(ctx *SymbolDeclarationContext) {
	key := ctx.IDENTIFIER().GetText()
	value := ctx.STRING_LITERAL().GetText()
	symbol := &Symbol{key, value}
	symbols = append(symbols, *symbol)
	symbolMaps[key] = value
}

func (s *DefineAppListener) EnterDefineDeclaration(ctx *DefineDeclarationContext) {
	switch ctx.DefineKey().GetText() {
	case "entry_point":
		buildEntryPoint(ctx)
	}
}

func buildEntryPoint(ctx *DefineDeclarationContext) {
	for _, express := range ctx.AllDefineExpress() {
		firstChild := express.(*DefineExpressContext).GetChild(0)
		typeOfExpress := reflect.TypeOf(firstChild).String()
		switch typeOfExpress {
		case "*parser.DefineAttributeContext":
			attributeCtx := firstChild.(*DefineAttributeContext)
			attributeCtx.GetText()
		case "*parser.DefineTemplateContext":
			defineCtx := firstChild.(*DefineTemplateContext)
			templateKey := defineCtx.IDENTIFIER().GetText()
			defineBody := defineCtx.DefineBody().(*DefineBodyContext)
			funcName := defineBody.STRING_LITERAL().GetText()
			templateData := defineBody.TemplateData().GetText()

			allSymbol := defineBody.AllSymbolKey()

			var codeText []string
			for index, symbolKey := range allSymbol {
				symbolCtx := symbolKey.(*SymbolKeyContext)
				symbolText := symbolCtx.GetText()

				codeText = append(codeText, symbolMaps[symbolText])
				if index == 0 {
					codeText = append(codeText, funcName)
				}
				if index == len(allSymbol)-2 {
					codeText = append(codeText, buildTemplate(templateData))
				}
			}

			defineInformation.DefineTemplates[templateKey] = strings.Join(codeText, " ")
		}
	}
}

func buildTemplate(templateData string) string {
	return startSymbol + templateData + endSymbol
}

func (s *DefineAppListener) getDefineInformation() DefineInformation {
	return defineInformation
}
