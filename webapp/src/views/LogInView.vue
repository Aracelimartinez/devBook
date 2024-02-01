<script setup>

</script>

<template>
  <main>
    <div id="login-page" class="h-dvh text-color-text-body">
      <div class="m-5 p-5 w-10/12">
        <img alt="LogIn image"  src="@/assets/undraw_right_places_re_3sve.svg"/>
      </div>
      <div class="w-10/12 p-4 flex flex-col text-base signin-card">
        <h1 class="my-3 text-xl font-bold">Iniciar sessão</h1>
        <input class="my-2 p-2 border rounded-md border-color-secondary hover:border-color-primary duration-150 fontAwesome" type="email" placeholder="&#xf0e0;  E-mail" >
        <input class="my-2 p-2 border rounded-md border-color-secondary hover:border-color-primary duration-150 fontAwesome" type="password" placeholder="&#xf084;  Senha" >
        <button class="mt-2 py-2 px-6  text-white bg-color-primary rounded-md hover:opacity-80 duration-150">Log in</button>
        <div class="my-2 flex justify-center text-line">
          <span>ou</span>
        </div>
        <button class="mb-3 py-2 px-6 bg-color-secondary rounded-md hover:opacity-80 duration-150">Criar conta</button>
      </div>
    </div>
  </main>
</template>

<style lang="scss" scoped>
  #login-page {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-gap: 16px;
    place-items: center;
  }

  .signin-card {
    background-color: #fff;
    box-shadow: 0px 0px 2px 2px rgba(0, 0, 0, 0.05);
    border-radius: 8px;
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
}

  // Smallest device
  @media (min-width: 100px) and (max-width: 575px) {
    #login-page {
      grid-template-columns: 1fr;
    }
  }

  // Small devices (landscape phones, 576px and up)
  @media (min-width: 576px) and (max-width: 767px){
    #login-page {
      grid-template-columns: 1fr;
    }
  }

  // Medium devices (tablets, 768px and up)
  @media (min-width: 768px) {
    #login-page {
      grid-template-columns: 1fr 1fr;
    }
  }
</style>
