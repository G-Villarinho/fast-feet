package validators

var ValidationMessages = map[string]string{
	"required": "Este campo é obrigatório. Por favor, preencha corretamente.",
	"email":    "O formato do e-mail está inválido. Certifique-se de que ele esteja no formato correto (exemplo@dominio.com).",
	"min":      "O valor informado é muito curto. Por favor, insira um valor com no mínimo {0} caracteres.",
	"max":      "O valor informado excede o limite máximo de {0} caracteres. Por favor, revise.",
	"eqfield":  "Os valores dos campos não coincidem. Verifique se ambos os campos foram preenchidos corretamente.",
	"gt":       "O valor informado deve ser maior que zero. Insira um valor válido.",
	"datetime": "O formato da data está incorreto. Por favor, use o formato válido (dd/mm/aaaa).",
	CPFTag:     "O formato do CPF está inválido. O formato correto é 999.999.999-99.",
}
