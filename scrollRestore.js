// ذخیره موقعیت اسکرول قبل از ترک صفحه
window.addEventListener('beforeunload', () => {
  sessionStorage.setItem('scrollPosition', window.scrollY);
});

// بازیابی موقعیت اسکرول هنگام بارگذاری صفحه
window.addEventListener('load', () => {
  const scrollPosition = sessionStorage.getItem('scrollPosition');
  if (scrollPosition !== null) {
    window.scrollTo(0, parseInt(scrollPosition));
  }
});

// برای صفحات SPA (Single Page Application)
window.addEventListener('pageshow', (event) => {
  if (event.persisted) {
    const scrollPosition = sessionStorage.getItem('scrollPosition');
    if (scrollPosition !== null) {
      window.scrollTo(0, parseInt(scrollPosition));
    }
  }
});
