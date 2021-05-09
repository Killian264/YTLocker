import React from 'react';
import { UserInfoBar } from '../components/UserInfoBar';

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

export const DashboardPage = ({ className }) => {

	return (
		<div className="m-4 mx-auto max-w-7xl" >
			<UserInfoBar className="flex-grow"  user={user} statCards={statsCards} ></UserInfoBar>
		</div>
	);
};

DashboardPage.propTypes = {};