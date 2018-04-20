import React from 'react';
import renderer from 'react-test-renderer';
import EmptyListCTA from './EmptyListCTA';

const model = {
  title: '标题',
  buttonIcon: 'ga css class',
  buttonLink: 'http://url/to/destination',
  buttonTitle: '点击我',
  proTip: '这是一个提示',
  proTipLink: 'http://url/to/tip/destination',
  proTipLinkTitle: '更多',
  proTipTarget: '_blank',
};

describe('EmptyListCTA', () => {
  it('renders correctly', () => {
    const tree = renderer.create(<EmptyListCTA model={model} />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
