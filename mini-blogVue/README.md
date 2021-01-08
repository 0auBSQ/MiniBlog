# mini-blog

Ce README détaille la partie front end du projet, je vais y expliqer
comment lancer Vue, les solutions techniques que nous avons choisi et pourquoi et enfin quelles ont été les dificultées.
Cette partie du projet utilise Node Js et npm.

## Project setup
```
npm install 
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Solutions techniques
### Vue
Pour le front end nous avons choisi d'utiliser Vue js pour les nombreux avantages qu'il procure, dans un premier temps c'est un framework Js qui permet un dévelopement plus simple et plus rapide, de plus, ce projet est voué à évoluer après la date de rendu, de ce fait, utiliser un frameworks Js qui décompose l'aplication en de nombreux composant indépendant rend la maintenance et l'amélioration de ce dernier beaucoup plus simple.

Avec Vue nous devons utiliser un certain nombre de librairies

## Vue X
Vue x permet la création d'un "store", ce store permet de centraliser des traitements, le store permet entre autre de faire communiquer les composants entre eux.

## Vue Router
Vue js permet de creer des single page aplication, si on veux bénéficier des url pour les différentes pages, comme par exemple acceder directement à un article en cherchant "http://localhost:8080/article/1" il nous faut utiliser Vue Router.
Pour Vue Router il suffit de créer une liste de route, d'y associer le composant que l'on souhaite et si besoin, faire des traitements relatif a l'url de départ et/ou à l'url d'arriver, nous permettant de limiter l'acces à certaine page, comme les page de compte par exemple.

## Axios
Nous avons utilisé Axios pour effectuer les requetes http

## SASS
Pour simplifier et accélérer le dévelopement des visuels de notre application nous avons utilisé le préprocesseur SASS

