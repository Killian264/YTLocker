import { Alert, AlertProps } from "../components/Alert";
import { sRadio, sString } from "./utils/utils";
import { Story } from "@storybook/react";
import { Dropdown } from "../components/Dropdown";

export default {
	title: "Alert",
	component: Alert,
};

const Mocked: Story<AlertProps & { message: string }> = ({ ...props }) => {
	return (
		<div>
			{/* <Dropdown
				className="ml-24"
				title="Account Selection"
				items={["Account settings", "Support", "License2", "Sign Out"]}
				OnItemSelected={() => "VOID"}
			></Dropdown> */}
			<Alert {...props}></Alert>
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	type: sRadio(["success", "failure"]),
	message: sString("This is an alert"),
};
