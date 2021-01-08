# mini-blog FR

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
La principale difficulté coté front-end à été de créer une application avec un bon UX et un bon UI, n'ayant jamais eu de véritable formation dans ce domaine nous allons faire évoluer ces points-ci, on peut ajouter à cela, AXIOS et son systeme de PROMISE.

## Les prochains ajouts
-Les prochaines modifications porteront surtout sur les bug visuel afin d'obtenir une application parfaitement stable.
-De régler les affichages correctement afin de corriger les bugs pour les appareils mobiles.
-D'envisager l'utilisation des technologies PWA.
-"j'aime" et "je n'aime pas" sur les commentaires
-Mise en place L10N


# mini-blog EN

This README details the front end part of the project, I will explain
how to launch Vue, the technical solutions that we chose and why and finally what were the difficulties.
This part of the project uses Node Js and npm.

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
For the front-end we chose to use Vue js for the many advantages it provides, first of all it is a Js framework which allows a simpler and faster development, moreover, this project is dedicated to evolving after the render date, therefore, using a Js framework that breaks down the application into many independent components makes maintaining and improving the latter much easier.

With Vue we have to use a number of libraries.

### Vue X
Vue x allows the creation of a "store", this store allows processing to be centralized, the store allows, among other things, to make the components communicate with each other.

### Vue Router
Vue js allows you to create Single Page Applications, if we want to benefit from the URLs for the different pages, for example to directly access an article by searching for "http://localhost:8080/article/1" we must use Vue Router .
For Vue Router, all you have to do is create a route list, associate it with the component you want and if necessary, perform processing relating to the starting URL and / or the arriving URL, allowing us to limit access to certain pages, such as account pages for example.

### Axios
We used Axios to make the http requests

### SASS
To simplify and accelerate the development of the visuals of our application we have used the SASS preprocessor

## Les dificultées
The main difficulty on the front-end side was to create an application with a good UX and a good UI, having never had any real training in this area, we are going to develop these points, we can add to that, AXIOS and its PROMISE system.

## Les prochains ajouts
-The next changes will mainly focus on visual bugs in order to obtain a perfectly stable application.
-To adjust the displays correctly in order to fix bugs for mobile devices.
-To consider the use of PWA technologies.
-Like and dislike comments
-L10N installation
