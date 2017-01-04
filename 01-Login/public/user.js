$(document).ready(function() {
    $('.btn-logout').click(function(e) {
      e.preventDefault();
      Cookies.remove('auth-session');
      window.location.href = '/';
    });
});
