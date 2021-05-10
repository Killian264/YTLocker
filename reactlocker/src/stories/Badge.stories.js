import React from "react";
import { Badge } from "../components/Badge";
import { toEnum } from "./utils/utils";

export default {
	title: "Badge",
	component: Badge,
	argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({ children, ...props }) => {
	return (
		<>
			<Badge {...props} className="mr-3" color={props.color || "primary"}>
				PRO
			</Badge>
			<Badge {...props} color={props.color || "primary"}>
				Color
			</Badge>
		</>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	color: toEnum(["primary", "secondary"]),
};
