import React from 'react';
import renderer from 'react-test-renderer';
import Tooltip from './Tooltip';

describe('Tooltip', () => {
  it('renders correctly', () => {
    const tree = renderer
      .create(
        <Tooltip className="test-class" placement="auto" content="Tooltip text">
          <a href="http://www.dataconnect.com">链接到提示</a>
        </Tooltip>
      )
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
