import { Story } from "@storybook/react";
import { LoginPage, LoginPageProps } from "../pages/LoginPage";

export default {
	title: "LoginPage",
	component: LoginPage,
};

const Mocked: Story<LoginPageProps> = ({ ...props }) => {
	return <LoginPage {...props}></LoginPage>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
