import React from 'react';

export interface IProps {
  password: string;
}

export class PasswordStrength extends React.Component<IProps, any> {

  constructor(props) {
    super(props);
  }

  render() {
    const { password } = this.props;
    let strengthText = "强度: 您的密码足够复杂.";
    let strengthClass = "password-strength-good";

    if (!password) {
      return null;
    }

    if (password.length <= 8) {
      strengthText = "强度: 您可以设置的更复杂一些.";
      strengthClass = "password-strength-ok";
    }

    if (password.length < 4) {
      strengthText = "强度: 您的密码太弱了.";
      strengthClass = "password-strength-bad";
    }

    return (
      <div className={`password-strength small ${strengthClass}`}>
        <em>{strengthText}</em>
      </div>
    );
  }
}


