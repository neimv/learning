AWSTemplateFormatVersion: "2010-09-09"

# Opcional
Description: Hello primer lambda (opcional)

# Opcional
# Metadata: opcional

# Valores que se le van a pasar a la plantilla
# Se le pueden enviar diferentes tipos de recursos
# envio de llaves, subredes
# se le pueden enviar una lista de direcciones
# se pueden usar expresiones regulares
Parameters:
  NombreLambda:
    Description: Nombre de la funcion lambda
    Type: String

  RuntimeLambda:
    Description: Ingresa el runtime de la funcion lambda
    Type: String
    Default: python3.7
    AllowedValues:
      - python3.7
      - python2.7
      - ruby2.5
      - nodejs8.10
      - java8
      - dotnetcore2.1

# Arreglo de llave valor referenciados en el template
# Se pueden especificar las regiones
Mappings:
  RegionMap:

