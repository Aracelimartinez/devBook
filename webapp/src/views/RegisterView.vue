<script setup>
import { useRouter } from 'vue-router';
import { useForm } from "vee-validate";
import { toTypedSchema } from '@vee-validate/yup';
import {string, object, ref } from 'yup';
import api from '@/services/api';

const router = useRouter();

const redirectToLogin = () => {
router.push('/login');
};

const { errors, handleSubmit, meta, defineField } = useForm({
  validationSchema: toTypedSchema( object({
      name: string().max(40).required(),
      nick: string().min(4).max(20).trim('Username cannot include leading and trailing spaces').strict().required(),
      email: string().email().required(),
      password: string().min(8).trim('Password cannot include leading and trailing spaces').strict().matches(
                /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])/,
                'Password must contain at least one number as well as one uppercase, lowercase, and special character'
            ).required(),
      passwordConfirm: string()
     .oneOf([ref('password')], 'passwords must match').required(),
    }),
  ),
});

const onSubmit = handleSubmit(async values => {
  const response = await api.post('/users', values);
  if (!response.errors) {
    console.log(response)
    return;
  }
  // set single field error
  if (response.errors) {
    console.log(response.errors)
    return;
    // setFieldError('email', response.errors.email);
  }
  // set multiple errors, assuming the keys are the names of the fields
  // and the key's value is the error message
  return;
  // setErrors(response.errors);
});

const [name, nameAttrs] = defineField('name', state => {
  return {
    validateOnModelUpdate: state.errors.length > 0,
  };
});

const [nick, nickAttrs] = defineField('nick', state => {
  return {
    validateOnModelUpdate: state.errors.length > 0,
  };
});

const [email, emailAttrs] = defineField('email', state => {
  return {
    validateOnModelUpdate: state.errors.length > 0,
  };
});

const [password, passwordAttrs] = defineField('password', state => {
  return {
    validateOnModelUpdate: state.errors.length > 0,
  };
});

const [passwordConfirm, passwordConfirmAttrs] = defineField('passwordConfirm', state => {
  return {
    validateOnModelUpdate: state.errors.length > 0,
  };
});
</script>

<template>
  <main>
    <div id="register-page" class="h-dvh text-color-text-body">

      <div class="w-10/12 flex flex-col">
        <img class="m-5 w-6/12" alt="LogIn image"  src="@/assets/images/undraw_fingerprint_login.svg"/>
        <h1 class="ml-5 mb-5 text-3xl font-bold text-left">Junte-se hoje à comunidade do <span class="text-color-primary">DevBook</span></h1>
      </div>

      <!-- Register card -->
      <div class=" flex flex-col p-7 text-base register-card w-10/12">
        <h2 class="mb-3 text-xl font-bold">Crie a sua conta!</h2>
        <form @submit="onSubmit" class="flex flex-col" name="register-fom">
          <input class="mt-5 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="text" placeholder="&#xf5b7;  Nome completo"  v-model="name" v-bind="nameAttrs">
          <span class="text-xs text-red-600">{{ errors.name }}</span>

          <input class="mt-5 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="text" placeholder="&#xf007;  Nome de usuário" v-model="nick" v-bind="nickAttrs">
          <span class="text-xs text-red-600">{{ errors.nick }}</span>

          <input class="mt-5 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="email" placeholder="&#xf0e0;  E-mail" v-model="email" v-bind="emailAttrs">
          <span class="text-xs text-red-600">{{ errors.email }}</span>

          <input class="mt-5 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="password" placeholder="&#xf084; Senha" v-model="password" v-bind="passwordAttrs">
          <span class="text-xs text-red-600">{{ errors.password }}</span>

          <input class="mt-5 p-2 border rounded-md border-color-details hover:border-color-primary duration-150 fontAwesome" type="password" placeholder="&#xf084;  Confirmar senha" v-model="passwordConfirm" v-bind="passwordConfirmAttrs">
          <span class="text-xs text-red-600">{{ errors.passwordConfirm }}</span>

          <button type="submit" class="mt-3 py-2 px-6  text-white bg-color-primary rounded-md hover:bg-color-secondary duration-150" :disabled="!meta.valid" >Registrar-se</button>
        </form>
        <div class="my-3 text-line">
          <span>ou</span>
        </div>
        <button @click="redirectToLogin" class="mb-3 py-2 px-6 bg-color-details rounded-md hover:opacity-80 duration-150">Faça seu login</button>
      </div>

    </div>
  </main>
</template>

<style lang="scss" scoped>
  #register-page {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 16px;
    place-items: center;
  }
  .register-card {
    background-color: #fff;
    box-shadow: 0px 0px 2px 2px rgba(0, 0, 0, 0.075);
    border-radius: 8px;
    max-width: 600px;
  }

  .text-line {
  position: relative;
  display: inline-block;
  text-align: center;
  }
  .text-line::before, .text-line::after {
    content: "";
    position: absolute;
    top: 0.65em;
    width: 40%;
    height: 0.05em;
    background-color: #666666;
    left: 0em;
  }
  .text-line::after {
    left: initial;
    right: 0em;
  }

  .fontAwesome {
  font-family: 'Roboto', FontAwesome, sans-serif;
  font-weight: 100;
  font-size: 14px;
  }

  // Smallest device
  @media (min-width: 100px) and (max-width: 575px) {
    #register-page {
      grid-template-columns: 1fr;
    }
  }

  // Small devices (landscape phones, 576px and up)
  @media (min-width: 576px) and (max-width: 767px){
    #register-page {
      grid-template-columns: 1fr;
    }
  }

  // Medium devices (tablets, 768px and up)
  @media (min-width: 768px) {
    #register-page {
      grid-template-columns: 1fr 1fr;
      margin: 0px 80px;
    }
  }
</style>
