import React from 'react';
import renderer from 'react-test-renderer';
import PickerOption from './PickerOption';

const model = {
  onSelect: () => {},
  onFocus: () => {},
  isFocused: () => {},
  option: {
    title: 'Model title',
    avatarUrl: 'url/to/avatar',
    label: '用户选取标签',
  },
  className: 'class-for-user-picker',
};

describe('PickerOption', () => {
  it('renders correctly', () => {
    const tree = renderer.create(<PickerOption {...model} />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
