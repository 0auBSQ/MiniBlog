import Vue from 'vue';
import { MLInstaller, MLCreate, MLanguage } from 'vue-multilanguage';

Vue.use(MLInstaller);

// Source :
// https://www.npmjs.com/package/vue-multilanguage

export default new MLCreate({
  initial: 'japanese',
  save: false,
  languages: [
    new MLanguage('english').create({
      home: 'Home',
      account: 'Account',
      login: 'Login',
      register: 'Register',
      search: 'Search',

      //title: 'Hello {0}!',
      //msg: 'You have {f} friends and {l} likes'
    }),

    new MLanguage('french').create({
      home: 'Accueil',
      account: 'Mon compte',
      login: 'Connection',
      register: 'S\'inscrire',
      search: 'Rechercher',

      //title: 'Oi {0}!',
      //msg: 'Você tem {f} amigos e {l} curtidas'
    }),

    new MLanguage('japanese').create({
      home: 'ホームページ',
      account: 'アカウント',
      login: 'ログイン',
      register: '登録する',
      search: '探る',

      //title: 'Oi {0}!',
      //msg: 'Você tem {f} amigos e {l} curtidas'
    })
  ],
  middleware: (component, path) => {
    const newPath = `${component.$options.name}.${path}`;
    return newPath;
  }
});
