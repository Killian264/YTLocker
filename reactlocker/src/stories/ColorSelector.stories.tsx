import { Story } from "@storybook/react";
import { useState } from "react";
import { ColorSelector, ColorSelectorProps } from "../components/ColorSelector";
import { Color } from "../shared/types";
import { sString } from "./utils/utils";

export default {
	title: "ColorSelector",
	component: ColorSelector,
};

const Mocked: Story<ColorSelectorProps & { message: string }> = ({ ...props }) => {
	let [selected, setSelected] = useState<Color>("red-1");

	return (
		<div>
			<ColorSelector
				disabled={["blue-1", "pink-1"]}
				selected={selected}
				onClick={(color) => {
					setSelected(color);
				}}
			></ColorSelector>
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
};
