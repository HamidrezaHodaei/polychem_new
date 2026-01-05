<template>
  <div class="map-wrapper">
    <div ref="mapContainer" class="map-container">
      <iframe
        width="100%"
        height="100%"
        style="border:0;"
        loading="lazy"
        allowfullscreen
        referrerpolicy="no-referrer-when-downgrade"
        :src="`https://maps.google.com/maps?q=${destLat},${destLng}&t=&z=13&ie=UTF8&iwloc=&output=embed`">
      </iframe>
    </div>
    
    <!-- Navigation Button -->
    <div class="routing-button">
      <button @click="getDirections" class="btn-directions">
        <span>Get Directions</span>
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const mapContainer = ref(null);

const destLat = 38.92111436229225;
const destLng = 45.6492204411566;

const getDirections = () => {
  const ua = navigator.userAgent || '';
  const isMobile = /Android|iPhone|iPad|iPod|Windows Phone|webOS/i.test(ua);
  const isIOS = /iPhone|iPad|iPod/i.test(ua) && !/Android/i.test(ua);
  const isAndroid = /Android/i.test(ua);

  const openUrl = (url) => {
    window.open(url, '_blank');
  };

  const buildWebMapsUrl = (userLat, userLng) => {
    const dest = encodeURIComponent(`${destLat},${destLng}`);
    if (typeof userLat === 'number' && typeof userLng === 'number') {
      const origin = encodeURIComponent(`${userLat},${userLng}`);
      return `https://www.google.com/maps/dir/?api=1&origin=${origin}&destination=${dest}&travelmode=driving`;
    }
    return `https://www.google.com/maps/dir/?api=1&destination=${dest}&travelmode=driving`;
  };

  // Mobile behavior: open native apps
  if (isMobile) {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const userLat = position.coords.latitude;
          const userLng = position.coords.longitude;

          if (isIOS) {
            const url = `maps://?saddr=${encodeURIComponent(userLat + ',' + userLng)}&daddr=${encodeURIComponent(destLat + ',' + destLng)}&dirflg=d`;
            openUrl(url);
            setTimeout(() => openUrl(buildWebMapsUrl(userLat, userLng)), 1000);
          } else if (isAndroid) {
            const intentUrl = `intent://maps/dir/?api=1&origin=${userLat},${userLng}&destination=${destLat},${destLng}&travelmode=driving#Intent;scheme=https;package=com.google.android.apps.maps;end`;
            openUrl(intentUrl);
            setTimeout(() => openUrl(buildWebMapsUrl(userLat, userLng)), 700);
          } else {
            openUrl(buildWebMapsUrl(userLat, userLng));
          }
        },
        (error) => {
          const destOnlyWeb = buildWebMapsUrl();
          if (isIOS) {
            openUrl(`maps://?daddr=${encodeURIComponent(destLat + ',' + destLng)}&dirflg=d`);
            setTimeout(() => openUrl(destOnlyWeb), 1000);
          } else if (isAndroid) {
            openUrl(`intent://maps/dir/?api=1&destination=${destLat},${destLng}&travelmode=driving#Intent;scheme=https;package=com.google.android.apps.maps;end`);
            setTimeout(() => openUrl(destOnlyWeb), 700);
          } else {
            openUrl(destOnlyWeb);
          }
          console.error('Geolocation error:', error);
        }
      );
    } else {
      const destOnlyWeb = buildWebMapsUrl();
      openUrl(destOnlyWeb);
    }
    return;
  }

  // Desktop: open Google Maps with directions
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(
      (position) => {
        const userLat = position.coords.latitude;
        const userLng = position.coords.longitude;
        openUrl(buildWebMapsUrl(userLat, userLng));
      },
      (error) => {
        openUrl(buildWebMapsUrl());
        console.error('Geolocation error:', error);
      }
    );
  } else {
    openUrl(buildWebMapsUrl());
  }
};
</script>

<style scoped>
.map-wrapper {
  width: 100%;
  height: 610px;
  position: relative;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  z-index: 1;
}

.map-container {
  width: 100%;
  height: 100%;
}

.routing-button {
  position: absolute;
  bottom: 15px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 10;
}

.btn-directions {
  position: relative;
  overflow: hidden;
  background: #f3f4f6;
  color: #848484;
  border: 1px solid #d1d5db;
  padding: 10px 20px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: color 300ms, border-color 300ms, background-color 300ms;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  white-space: nowrap;
  z-index: 1;
}

.btn-directions::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: #FFCD05;
  transform: translateY(-100%);
  transition: transform 300ms ease;
  z-index: 0;
}

.btn-directions:hover::before,
.btn-directions:focus-visible::before {
  transform: translateY(0);
}

.btn-directions:hover,
.btn-directions:focus-visible {
  color: #fff;
  border-color: #FFCD05;
  outline: none;
}

.btn-directions span {
  position: relative;
  z-index: 1;
}
</style>