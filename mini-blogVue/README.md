# mini-blog

Ce README détaille la partie front end du projet, je vais y expliquer
comment lancer Vue, les solutions techniques que nous avons choisies et pourquoi et enfin quelles ont été les dificultés.
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

## Solutions techniques
### Vue
Pour le front-end nous avons choisi d'utiliser Vue js pour les nombreux avantages qu'il procure, dans un premier temps c'est un framework Js qui permet un développement plus simple et plus rapide, de plus, ce projet est voué à évoluer après la date de rendu, de ce fait, utiliser un framework Js qui décompose l'application en de nombreux composants indépendants rend la maintenance et l'amélioration de ce dernier beaucoup plus simple.

Avec Vue nous devons utiliser un certain nombre de librairies.

### Vue X
Vue x permet la création d'un "store", ce store permet de centraliser des traitements, le store permet entre autres de faire communiquer les composants entre eux.

### Vue Router
Vue js permet de créer des Single Page Application, si on veut bénéficier des URL pour les différentes pages, comme par exemple accéder directement à un article en cherchant "http://localhost:8080/article/1" il nous faut utiliser Vue Router.
Pour Vue Router il suffit de créer une liste de route, d'y associer le composant que l'on souhaite et si besoin, faire des traitements relatifs à l'URL de départ et/ou à l'URL d'arriver, nous permettant de limiter l'accès à certaine pages, comme les pages de compte par exemple.

### Axios
Nous avons utilisé Axios pour effectuer les requêtes http

### SASS
Pour simplifier et accélérer le développement des visuels de notre application nous avons utilisé le préprocesseur SASS

## Les dificultées
La principale difficulté coté front-end à été de créer une application avec un bon UX et un bon UI, n'ayant jamais eu de véritable formation dans ce domaine nous allons faire évoluer ces points-ci.

## Les prochains ajouts
-Les prochaines modifications porteront surtout sur les bug visuel afin d'obtenir une application parfaitement stable.
-De régler les affichages correctement afin de corriger les bugs pour les appareils mobiles.
-D'envisager l'utilisation des technologies PWA.
