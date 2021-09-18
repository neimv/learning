#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Sun Jul 11 13:59:28 2021

@author: neimv
"""

import collections
import math
import random

from bokeh.plotting import figure, show

# Programación dínamica
# - en los 50's Richard Bellman
# - se esconde para que no se sepa que haga matemáticas
# - no tiene que ver con dinamismo en programación
# - Subestructura óptima: solucion global optima se pueden encontrar usando soluciones optimas de subproblemas locales
# - problemas empalmados: una solución óptima que involucra resolver el mismo problema varias ocasiones
# Optimización (Memoization) 
# - tecnica para guardar resultados previos y evitar hacerlos de nuevo
# - normalmente se usa un diccionario, donde las consultas son en O(1)
# - intercambio tiempo por espacio

# Optimización de fibionacci (F_n = F_n-1 + F_n-2)
def fibionacci_recursivo(n):
    if n == 0 or n == 1:
        return 1
    
    return fibionacci_recursivo(n - 1) + fibionacci_recursivo(n - 2)


def fibionacci_dinamico(n, memo={}):
    if n == 0 or n == 1:
        return 1

    try:
        return memo[n]
    except KeyError:
        resultado = fibionacci_dinamico(n - 1, memo) + fibionacci_dinamico(n - 2, memo)
        memo[n] = resultado
        
        return resultado

# Caminos aleatorios
# - es un tipo de simulación que elige aleatoriamente una desición dentro de un conjunto de decisiones validas
# - Se utiliza en muchos campos del conocimiento cuando los sistemas no son deterministras e incluyen elementos de aleatoriedad
# - no se pueden usar programas deterministicos
# - simulación de particulas, no solo es a nivel microscopico, también macroscopico
# - se puede usar en acciones (dineros)
# Camino de borrachos

class Borracho:
    def __init__(self, nombre):
        self.nombre = nombre


class BorrachoTradicional(Borracho):
    def __init__(self, nombre):
        super().__init__(nombre)
        
    def camina(self):
        return random.choice([(0, 1), (0, -1), (1, 0), (-1, 0)])
    
    
class Coordenada:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        
    def mover(self, delta_x, delta_y):
        return Coordenada(self.x + delta_x, self.y + delta_y)
    
    def distancia(self, otra_coordenada):
        delta_x = self.x - otra_coordenada.x
        delta_y = self.y - otra_coordenada.y
        
        return (delta_x**2 + delta_y**2)**0.5
    

class Campo:
    def __init__(self):
        self.coordenadas_de_borrachos = {}
        
    def anadir_borracho(self, borracho, coordenada):
        self.coordenadas_de_borrachos[borracho] = coordenada
        
    def mover_borracho(self, borracho):
        delta_x, delta_y = borracho.camina()
        coordenada_actual = self.coordenadas_de_borrachos[borracho]
        nueva_coordenada = coordenada_actual.mover(delta_x, delta_y)
        
        self.coordenadas_de_borrachos[borracho] = nueva_coordenada

    def obtener_coordenada(self, borracho):
        return self.coordenadas_de_borrachos[borracho]
    

#### Camino aleatorio
def caminata(campo, borracho, pasos):
    inicio = campo.obtener_coordenada(borracho)
    
    for _ in range(pasos):
        campo.mover_borracho(borracho)
        
    return inicio.distancia(campo.obtener_coordenada(borracho))


def simular_caminata(pasos, numero_de_intentos, tipo_de_borracho):
    borracho = tipo_de_borracho(nombre='Hello')
    origen = Coordenada(0, 0)
    distancias = []
    
    for _ in range(numero_de_intentos):
        campo = Campo()
        campo.anadir_borracho(borracho, origen)
        simulacion_caminata = caminata(campo, borracho, pasos)
        distancias.append(round(simulacion_caminata, 1))
        
    return distancias


def graficar(x, y):
    grafica = figure(title="Camino aleatorio", x_axis_label='pasos', y_axis_label='distancia recorrida')
    grafica.line(x, y, legend='distancia media')
    
    show(grafica)


def main(distancias_de_caminata, numero_de_intentos, tipo_de_borracho):
    distancias_media_por_caminata = []    

    for pasos in distancias_de_caminata:
        distancias = simular_caminata(pasos, numero_de_intentos, tipo_de_borracho)
        distancia_media = round(sum(distancias) / len(distancias), 4)
        distancia_maxima = max(distancias)
        distancia_minima = min(distancias)
        distancias_media_por_caminata.append(distancia_media)
        
        print(f'{tipo_de_borracho.__name__} caminata aleatoria de {pasos} pasos')
        print(f'Distancia media: {distancia_media}')
        print(f'Distancia maxima: {distancia_maxima}')
        print(f'Distancia minima: {distancia_minima}')

    graficar(distancias_de_caminata, distancias_media_por_caminata)

# Programación estocástica
# - un programa es deterministico si el input produce el mismo output
# - los programas deterministicos son importantes pero no todos los problemas se pueden resolver asi
# - la programacion estocastica nos permite meter aleatoriedad a situacion de problmas mas complejos
# - aprovechan las distribuciones probabilisticas
# - pueden ser definidas si se conoce la poblacion
# - si no estan definidas se pueden calcular por estadistica inferencial
# Calculo de probabilidades
# - es una medida de la certidumbre de que un evento suceda o no suceda, un rango entre 0 y 1
# - 0 no sucede el evento
# - 1 es garantia de que suceda
# - se pregunta que fraccion de los eventos que van a suceder
# - se deben de poder calcular todos los eventos posibles
# - ley del complemento: sucede + no sucede = 1
# - probabilidad multiplicativa: P(A y B): P(A) * P(B)
# - aditiva, mutuamente exclusivos: P(A o B): P(A) + P(B)
# - aditiva, mutuamente no exclusivos: P(A o B): P(A) + P(B) - P(A y B)
# Simulación de probabilidades
def tirar_dado(numero_de_tiros):
    secuencia_de_tiros = []
    
    for _ in range(numero_de_tiros):
        tiro = random.choice(range(1, 7))
        secuencia_de_tiros.append(tiro)
    
    return secuencia_de_tiros


def main2(numero_de_tiros, numero_de_intentos):
    tiros = []
    
    for _ in range(numero_de_intentos):
        secuencia_de_tiros = tirar_dado(numero_de_tiros)
        tiros.append(secuencia_de_tiros)

    tiros_con_1 = 0
    
    for tiro in tiros:
        if 1 in tiro:
            tiros_con_1 += 1
            
    probabilidad_tiros_con_1 = tiros_con_1 / numero_de_intentos
    
    print(f'Probabilidad de obtener por lo menos un 1 en {numero_de_tiros} tiros = {probabilidad_tiros_con_1}')

# Inferencia estadistica
# - con simulaciones se pueden calcular las probabilidad de eventos complejos con probabilidades simples
# - si no se tiene las probalidades simples se puede hacer uso de tecnicas de estadistica inferencial, donde a partir de una muestra aleatoria se pueden saber sus propiedades
# - una muestra aleatoria muestra las mismas propiedades que la población de la cual fue extraida
# - ley de los grandes numeros: en pruebas independientes repetidas con la misma probabilidad p, la fraccioón de desviaciones de p converge a cero conforme la cantidad de muestreos sea infinita
# - falacia del apostador
#   . despues de un evento extremo ocurriran eventos menos extremos para nivelar la media, no sucede, los eventos son independientes
#   . despues de un evento extremo probablemente sea menos extremo, pero no nivela
# - Media, es el promedio de los valores, medida de tendencia central, de la poblacion es Mu y la de la muestra es x_bar (osea arriba una rayita de la x)
# - Varianza y desviación estandar: medidas de propagacion, que tan alejados o cerca estan de la media
#   . varianza: siempre se hace en relación de la media, que tan dispersos estan
#   . desviación estandar: raiz cuadrada de la varianza, variabilidad de las unidades que se usa
def media(X):
    return sum(X) / len(X)


def varianza(X):
    mu = media(X)
    
    acumulador = 0
    
    for x in X:
        acumulador += (x - mu)**2
        
    return acumulador / len(X)


def desviacion_estandar(X):
    return math.sqrt(varianza(X))


# Distribucion normal
# - es de las mas comunes
# - se define por su media y su desviacion estandar
# - permite calcular intervalos de confianza
# - regla empirica: 68-95-99.7: como se distribuye adentro de una distribucion normal los datos
# - cual es la dispersión de los datos a uno, dos o tres sigmas
# - permite calcular probabilidades con la densidad de la distribución normal
# Montecarlo
# - calculo de la probabilidad del juego perfecto del solitario
# - se simulan los juegos de montecarlo, juegos de casino de azar
# - permite crear simualciones para predecir el resultado
# - convierte problemas deterministicos en problemas estocasticos
# - se utiliza en muchas areas, desde la ingenieria hasta la biologia y el derecho
PALOS = ['espada', 'corazon', 'rombo', 'trebol']
VALORES = [
    'as', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'jota', 'reina', 'rey'
]

def crear_baraja():
    barajas = []
    
    for palo in PALOS:
        for valor in VALORES:
            barajas.append((palo, valor))
            
    return barajas


def obtener_mano(barajas, tamano_mano):
    mano = random.sample(barajas, tamano_mano)
    
    return mano


def main3(tamano_mano, intentos):
    barajas = crear_baraja()
    
    manos = []
    
    for _ in range(intentos):
        mano = obtener_mano(barajas, tamano_mano)
        manos.append(mano)
        
    pares = 0
    for mano in manos:
        valores = []
        for carta in mano:
            valores.append(carta[1])
            
        counter = dict(collections.Counter(valores))
        
        for val in counter.values():
            if val == 2:
                pares += 1
                break
            
    probabilidad_par = pares / intentos
    print(f"La probabilidad de obtener un par en una mano de {tamano_mano} es {probabilidad_par}")


# Cálculo de Pi
def aventar_agujas(numero_de_agujas):
    adentro_del_circulo = 0
    
    for _ in range(numero_de_agujas):
        x = random.random() * random.choice([-1, 1])
        y = random.random() * random.choice([-1, 1])
        distancia_desde_el_centro = math.sqrt(x ** 2 + y ** 2)
        
        if distancia_desde_el_centro <= 1:
            adentro_del_circulo += 1
            
    return (4 * adentro_del_circulo) / numero_de_agujas


def estimacion(numero_de_agujas, numero_de_intentos):
    estimados = []
    
    for _ in range(numero_de_intentos):
        estimacion_pi = aventar_agujas(numero_de_agujas)
        estimados.append(estimacion_pi)
        
    media_estimados = media(estimados)
    sigma = desviacion_estandar(estimados)
    
    print(f'Est = {round(media_estimados, 5)}, sigma={round(sigma, 5)}, agujas={numero_de_agujas}')

    return media_estimados, sigma


def estimar_pi(precision, numero_de_intentos):
    numero_de_agujas = 1000
    sigma = precision
    
    while sigma >= precision / 1.96:
        media, sigma = estimacion(numero_de_agujas, numero_de_intentos)
        numero_de_agujas *= 2
        
    return media

# Muestreo
# - comparación entre un ejemplo con otros
# - es por la limitación del hardware
# - las muestras tienden a tener las mismas propiedades que la poblacion total
# - el muestreo probabilistico
#   . muestreo aleatorio, cualquier miembro de la población tiene la misma probabilidad de ser escogido
#   . estratificado, se toman en cuenta las caracteristicas de la población para partirla en subgrupos, se incremente la probabilidad de que el muestreo sea representativo
# Teorema del límite central
# - de los mas importantes
# - nos ayuda a transformar cualquier distribución en una distribución normal
# - se toma cualquier distribucion, se obtiene una muestra, se le obtiene la media, la distribucion de la media que se obtengan terminara siendo normal
# Trabajar con datos experimentales
# - la aplicacion del metodo cientifico
# - se comienza con una teoria o hipotesis
# - se deben crear experimentos para falsear o validar esa hipotesis
# - se miden las diferencias entre las mediciones experimentales y aquellas predichas
# Regresión Líneal
# - permite aproximar una función a un conjkunto de datos obtenidos de manera experimental
# - no necesarimante permite aproximar funciones lineales, puede ser cualquier función polinomica


if __name__ == '__main__':
    # n = int(input('Escoge un número: '))
    # resultado = fibionacci_dinamico(n)
    # print(resultado)
    # distancias_de_caminata = [10, 100, 1000, 10000]
    # numero_de_intentos = 100

    # main(distancias_de_caminata, numero_de_intentos, BorrachoTradicional)
    # numero_de_tiros = int(input('Cuantos tiros del dado: '))
    # numero_de_intentos = int(input('Cuantas veces correra la simulación: '))
    
    # main2(numero_de_tiros, numero_de_intentos)
    # X = [random.randint(1, 21) for i in range(20)]
    # mu = media(X)
    # Var = varianza(X)
    # sigma = desviacion_estandar(X)

    # print(f'Arreglo X: {X}')
    # print(f'Media = {mu}')
    # print(f'varianza: {Var}')
    # print(f'desviacion estandar: {sigma}')
    # barajas = crear_baraja()
    # mano = obtener_mano(barajas, 5)
    # print(mano)
    # tamano_mano = int(input("de cuantas barajas sera la mano: "))
    # intentos = int(input("cuantos intentos para calcular la probabilidad: "))
    # main3(tamano_mano, intentos)
    estimar_pi(0.01, 1000)
