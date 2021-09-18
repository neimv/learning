#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Fri Jul 16 09:58:16 2021

@author: neimv
"""

# Pensamiento probabilistico
# - calcular las probabilidades de lo que pasa al rededor
# - como actualizar las probabilidades
# - la programacion probabilistica, se usa en varias áreas
# - se obtiene evidencia para saber si las hipotesis son correctas o falsas
# - se usan modelos probabilisticos
# - existen lenguajes y librerias para este tipo de problemas
# Probabilidades condicionales
# - cuantas veces sucede un evento dentro de la totalidad de eventos posibles
# - P(A & B) = P(A)P(B|A)
# - P(B) = P(A) * P(B|A) + P(not A)P(B|not A)
# Teorema de Bayes
# - P(A|B) = (P(B|A)P(A)) / P(B)
# - Prior: conocimiento previo de cierta situacion, hipotesis que tenemos
# - Posterior: una vez tenida la evidencia como actualizar las creencias
# - ecosograma: rectangulo donde se puede determinar las seccion para interpertar directamente los nombres
# - likelihood: que tan ciertos estamos en que la situacion es correcta
# Analisis de sintomas
def calc_bayes(prior_A, prob_B_dado_A, prob_B):
    return (prior_A * prob_B_dado_A) / prob_B


# Mentiras estadísticas
# - Garbage in, garbage out (GIGO)
#   . si avientas basura escupe basura
#   . la calidad de los datos es importante igual que la precisión de los computos
#   . si tenemos datos errados nuestros resultados seran erroneos
#   . Motor analitico de babach
# - Imagenes engañosas
#   . la estadistica se pueden usar para mentir... llegar a conclusiones inexactas
#   . es muy comun que se trate de engañar así
#   . generalmente pasa cuando se juega con escalas
#   . sin etiquetas no se debe confiar
#   . muy normal en noticieros
# Cum hoc Ergo Propter Hoc
#   . la correlacion entre dos variables no es casualidad, solo se mueven en el mismo sentido
#   . de -1 a 1, nos ayuda a saber com estan relacionadas
#   . despues de esto, eso; entonces a consecuencia de esto, eso
#   . el movimiento de un molino depende de muchas variables no solo del viento
#   . para evitarlo se debé pensar "out of the box"
# Prejucio en el muestreo
#   . si no se tiene aleatoriedad o representatividad
#   . muy comun en ciencias sociales...
#   . resultados de estudiantes universitarios, ya que generalmente encuentras mas estudiantes de ciencias sociales
#   . los comun es pensar que los antepasdos todos eran hombres de cavernas
# Falacia del francotirador de texas
#   . no tomar la aleatoriedad en consideracion
#   . cuando uno se enfoca demasiado en las similutudes y no se tiene una muestra amplia para revisar eso
#   . se falla en una hipotesis antes de recolectar datos, primero hipotesis luego los datos
#   . pasa mucho en los emprendedores donde siempre toman lo que les funciono vs lo que no funciono
# Porcentajes confusos
#   . es una de las tecnicas mas comunes para mentir con numeros
#   . cuando no se sabe de donde estan naciendo estos porcentajes
#   . siempre es importante ver el contexto
#   . los porcentajes en vacio no significan mucho
#   . se podría decir comparado con que
# Falacia de regresión
#   . muchos eventos fluctuan naturalmente... regresion a la media, los eventos extremos
#   . cuando se aplican medidas correctivas, se cae en este error, los cambios random en eventos, no sabemos que tiene que ver
#   . es muy comun en personas que se basan mas en castigar que en premiar
# # Machine learning
# - es la capacidad de las computadores de aprender sin ser explicitamente programadas
# - no hay inteligencia a nivel evolutivo
# - los humanos memorizados y generalizamos
# - se usa cuando programar un algoritmo es imposible
# - si el problema es muy complejo o no se conocer algoritmos para resolverlo
# - sirve para entender patrones
# - existe: aprendizaje supervisado, no supervisado y semisupervisado
# - batch y online learning
# Feature vectors
# - se usan para representar carateristicas simbolicas o numericas llamadas features, datos matematicos
# - permite analizar un objeto desde una perspectiva matematica
# - los algoritmos requieren representaciones numericas para poder ejecutar el computo
# - uno de los mas comunes es la presentacion del color a travez del RGB
# Metricas de distancia
# - muchos algoritmos pueden clasificarse como algoritmos de optimizacion
# - se desea optimizar la distancia entre vectores
# - manhattan solo permite ir entre calles, no directamente como la euclidiana
# Agrupamiento
# - Clustering:
#   . nos permiten entender la estructura de los datos, cuando no hay etiquetas
#   . permiten agrupar en diferentes grupos (clusters), para saber que tan diferentes son o iguales
# - Agrupamiento jerarquico
#   . algoritmo sencillo
#   . tomar los puntos mas cercanos, los agrupa en cluster y revisa las distancias y agrupa, se hace de forma iterrativa
#   . se debe pensar en las relaciones que se tienen en cada grupo
#   . parte en que cada datapoint es un grupo individual, checa las distancias entre cada punto y hace nuevos grupos
#   . produce un objeto que se llama dendograma, que sirve para ver que relaciones tienen los puntos
# - Kmeans
#   . agrupa haciendo centroides al azar
#   . k define el numero de grupos que se quiere hacer
#   . en cada iteracion el punto se ajusta a su centroide
# Clasificación
# - proceso mediante el cual se prodecie la clase de cierto dato
# - es un tipo de aprendizaje supervisado, se necesitan labels
# - se usa en muchos dominios
# - sigue los pasos: aprendizaje y clasificacion
# - K nearest neighbors
#   . parte de que se tiene una clasificacion previa
#   . encontrar la clasificacion de un nuevo datapoint
#   . K numero de vecinos para saber como clasificar
#   . es sencillo de implementar
#   . computacionalmente muy costoso
# ERRORES: 
# Antes de que se realice un examen médico, ¿cuál es tu probabilidad (sin mayor información) de padecer una enfermedad?
# ¿Cuál de las siguientes NO es una aplicación del Teorema de Bayes?
# El centroide, dentro del agrupamiento de K-means, ¿es estático?
# ¿Cuál NO es un algoritmo de agrupamiento?



if __name__ == '__main__':
    prob_cancer = 1 / 100000
    prob_sintoma_dado_cancer = 1
    prob_sintoma_dado_no_cancer = 10 / 99999
    prob_no_cancer = 1 - prob_cancer
    
    prob_sintoma = (
        prob_sintoma_dado_cancer * prob_cancer
    ) + (prob_sintoma_dado_no_cancer * prob_no_cancer)
    
    prob_cancer_dado_sintoma = calc_bayes(
        prob_cancer, prob_sintoma_dado_cancer, prob_sintoma
    )
    print(prob_cancer_dado_sintoma)
