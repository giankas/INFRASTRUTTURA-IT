var app = angular.module('myApp', []);

app.controller('MainController', function($scope, $http) {
  $scope.view = "login";
  $scope.isAuthenticated = false;
  $scope.credentials = {};
  $scope.regData = {};
  $scope.tickets = [];
  $scope.services = [];
  $scope.packages = [];
  $scope.orders = [];
  $scope.newTicket = {};
  $scope.contact = {};

  // Imposta la view corrente
  $scope.setView = function(viewName) {
    $scope.view = viewName;
    if(viewName === "dashboard" && $scope.isAuthenticated) {
      $scope.getTickets();
      $scope.getServices();
    }
    if(viewName === "packages" && $scope.isAuthenticated) {
      $scope.getPackages();
      $scope.getOrders();
    }
  };

  // Logout
  $scope.logout = function() {
    $scope.isAuthenticated = false;
    $scope.credentials = {};
    delete $http.defaults.headers.common.Authorization;
    $scope.setView("login");
  };

  // Login
  $scope.login = function() {
    $http.post('/api/login', $scope.credentials).then(function(response) {
      $scope.token = response.data.token;
      $scope.isAuthenticated = true;
      $http.defaults.headers.common.Authorization = 'Bearer ' + $scope.token;
      $scope.setView("dashboard");
    }, function(error) {
      alert('Credenziali non valide');
    });
  };

  // Registrazione
  $scope.register = function() {
    $http.post('/api/register', $scope.regData).then(function(response) {
      alert(response.data.message);
      $scope.setView("login");
    }, function(error) {
      alert(error.data.error || 'Errore durante la registrazione');
    });
  };

  // Recupera i ticket
  $scope.getTickets = function() {
    $http.get('/api/tickets').then(function(response) {
      $scope.tickets = response.data;
    });
  };

  // Crea un ticket
  $scope.createTicket = function() {
    $http.post('/api/tickets', $scope.newTicket).then(function(response) {
      $scope.tickets.push(response.data);
      $scope.newTicket = {};
    });
  };

  // Recupera i servizi attivi
  $scope.getServices = function() {
    $http.get('/api/services').then(function(response) {
      $scope.services = response.data;
    });
  };

  // Recupera i pacchetti
  $scope.getPackages = function() {
    $http.get('/api/packages').then(function(response) {
      $scope.packages = response.data;
    });
  };

  // Simula l'acquisto di un pacchetto
  $scope.buyPackage = function(packageId) {
    var order = { package_id: packageId };
    $http.post('/api/orders', order).then(function(response) {
      alert('Ordine creato! ID ordine: ' + response.data.id);
      $scope.getOrders();
    });
  };

  // Recupera gli ordini dell'utente
  $scope.getOrders = function() {
    $http.get('/api/orders').then(function(response) {
      $scope.orders = response.data;
    });
  };

  // Invio del form contatto
  $scope.sendContact = function() {
    $http.post('/api/contact', $scope.contact).then(function(response) {
      alert(response.data.message);
      $scope.contact = {};
    });
  };
});
