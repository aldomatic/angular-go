(function() {
  'use strict';

  angular
    .module('angularGo')
    .controller('MainController', MainController);

  /** @ngInject */
  function MainController($timeout) {
    var vm = this;
    vm.title = "Angular and Go!!"
    activate();


    function activate() {
      $timeout(function() {
        console.log("active()")
      }, 3000);
    }
  }
})();
