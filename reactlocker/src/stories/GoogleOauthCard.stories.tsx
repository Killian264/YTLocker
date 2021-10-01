import { Story } from "@storybook/react";
import { GoogleOAuthCard, GoogleOAuthCardProps } from "../components/GoogleOauthCard";
import { sRadio } from "./utils/utils";

export default {
	title: "GoogleOAuthCard",
	component: GoogleOAuthCard,
};

const Mocked: Story<GoogleOAuthCardProps> = ({ ...props }) => {
	return (
		<GoogleOAuthCard
			{...props}
		></GoogleOAuthCard>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	status: sRadio(["login", "link"]),
};