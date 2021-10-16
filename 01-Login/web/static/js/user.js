$(document).ready(function() {
    $('.btn-logout').click(function(e) {
      Cookies.remove('auth-session');
    });
});
