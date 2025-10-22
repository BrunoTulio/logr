# TODO: Criar Adaptador Logrus para Logr

## Passos do Plano

- [x] Criar diretório adapters/logrus.v1/
- [x] Criar logger.go: Implementar struct logger com *logrus.Logger incorporado, satisfazendo logr.Logger
- [x] Criar field.go: Funções para converter logr.Field para logrus.Fields
- [x] Criar level.go: Mapear níveis string para logrus.Level
- [x] Criar option.go: Definir Option struct e FnOption para configuração
- [x] Verificar compilação e integração com logr
- [x] Adicionar exemplo em examples/ para demonstração (opcional)
