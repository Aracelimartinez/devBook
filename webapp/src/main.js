import './assets/tailwind.css'

import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
// import { Field, Form, ErrorMessage, useField } from 'vee-validate';
// import '@vee-validate/rules'; //

const app = createApp(App)

app.use(router)
// app.use(Field);
// app.use(Form);
// app.use(ErrorMessage);
// app.use(useField);

app.mount('#app')
