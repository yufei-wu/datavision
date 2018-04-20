import coreModule from '../core_module';

export class ResetPasswordCtrl {
  /** @ngInject */
  constructor($scope, contextSrv, backendSrv, $location) {
    contextSrv.sidemenu = false;
    $scope.formModel = {};
    $scope.mode = 'send';

    var params = $location.search();
    if (params.code) {
      $scope.mode = 'reset';
      $scope.formModel.code = params.code;
    }

    $scope.navModel = {
      main: {
        icon: 'gicon gicon-branding',
        text: '重设密码',
        subTitle: '重设您的DataVision密码',
        breadcrumbs: [{ title: '登录', url: 'login' }],
      },
    };

    $scope.sendResetEmail = function() {
      if (!$scope.sendResetForm.$valid) {
        return;
      }
      backendSrv.post('/api/user/password/send-reset-email', $scope.formModel).then(function() {
        $scope.mode = 'email-sent';
      });
    };

    $scope.submitReset = function() {
      if (!$scope.resetForm.$valid) {
        return;
      }

      if ($scope.formModel.newPassword !== $scope.formModel.confirmPassword) {
        $scope.appEvent('alert-warning', ['新密码不匹配', '']);
        return;
      }

      backendSrv.post('/api/user/password/reset', $scope.formModel).then(function() {
        $location.path('login');
      });
    };
  }
}

coreModule.controller('ResetPasswordCtrl', ResetPasswordCtrl);
