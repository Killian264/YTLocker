import React from 'react';
import { UserInfoBar } from '../components/UserInfoBar';

export default {
  title: 'UserInfoBar',
  component: UserInfoBar,
  argTypes: { onClick: { action: "clicked" } },
};

const user = {
  username: "Killian",
  email: "killiandebacker@gmail.com",
  joined: "Mar 13 2021",
}

const statsCards = [
  {
    header: "Playlists",
		count: 454,
		measurement: "total",
  },
  {
    header: "Videos", 
    count: 357, 
    measurement: "total"
  },
  {
    header: "Subscriptions",
		count: 17,
		measurement: "total",
  },
  {
    header: "Updated",
		count: 13,
		measurement: "seconds ago",
  },

]

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<UserInfoBar user={user} statCards={statsCards} ></UserInfoBar>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {
	
};