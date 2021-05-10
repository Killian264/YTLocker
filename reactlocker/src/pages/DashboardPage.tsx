import React from "react";
import { UserInfoBar } from "../components/UserInfoBar";
import { StatCard } from "../shared/types";

const user = {
	username: "Killian",
	email: "killiandebacker@gmail.com",
	joined: "Mar 13 2021",
};

const stats: StatCard[] = [
	{
		header: "Playlists",
		count: 454,
		measurement: "total",
	},
	{
		header: "Videos",
		count: 357,
		measurement: "total",
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
];

export const DashboardPage: React.FC<{}> = () => {
	return (
		<div className="p-4 mx-auto max-w-7xl">
			<UserInfoBar
				className="flex-grow"
				user={user}
				stats={stats}
			></UserInfoBar>
		</div>
	);
};

DashboardPage.propTypes = {};
