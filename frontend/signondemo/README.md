# signondemo
## 路由
```js
  routes: [
      {
        path: '/signin',
        name: 'signin',
        component: () => import('../views/SignInView.vue')
      },
      {
        path: '/signup',
        name: 'signup',
        component: () => import('../views/SignUpView.vue')
      },
      {
        path: '/userhome',
        name: 'userhome',
        component: () => import('../views/UserHomeView.vue')
      }
    ]
```
## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
