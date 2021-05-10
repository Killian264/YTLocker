import { Story } from "@storybook/react";
import { Login, LoginProps } from "../components/Login";

export default {
	title: "Login",
	component: Login,
};

const Mocked: Story<LoginProps> = ({ ...props }) => {
	return (
		<Login
			{...props}
			onSubmit={(user) => console.log("Submitted", user)}
			onClickRegister={() => console.log("Swapped to register page.")}
		></Login>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
