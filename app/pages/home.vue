<template>
  <div>
    <Navbar/>
    <Hero2/>
    <div class="mt-8 slider">
      <Slider/>
      <Text/>
      <Cards/>
      <global/>
      <infinityslider/>
    </div>
    <footer1/> 
  </div>
</template>

<script setup>
  import Navbar from "~/components/navbar.vue"
  import Slider from "~/components/Slider.vue"
  import Hero from "~/components/Hero2.vue"
  import Cards from "~/components/cards.vue"
  import Text from "~/components/Text.vue"
  import global from "~/components/world.vue"
  import infinityslider from "~/components/InfiniteSlider.vue"
  import footer1 from "~/components/footer.vue";

</script> 


