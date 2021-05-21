import { Story } from "@storybook/react";
import { LoadingUserInfoBar } from "../components/LoadingUserInfoBar";
import { UserInfoBar, UserInfoBarProps } from "../components/UserInfoBar";

export default {
	title: "UserInfoBar",
	component: UserInfoBar,
};

const user = {
	username: "Killian",
	email: "killiandebacker@gmail.com",
	joined: new Date(),
};

const statsCards = [
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

const Mocked: Story<UserInfoBarProps> = ({ ...props }) => {
	return (
		<>
			<UserInfoBar user={user} stats={statsCards}></UserInfoBar>
			<LoadingUserInfoBar></LoadingUserInfoBar>
		</>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
