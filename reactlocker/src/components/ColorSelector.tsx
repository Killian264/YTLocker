import React from "react";
import { ColorArray } from "../shared/colors";
import { Color } from "../shared/types";
import { ColorBadge } from "./ColorBadge";

export interface ColorSelectorProps {
	className?: string;
	selected: Color | null;
	disabled: Color[];
	onClick: (color: Color) => void;
}

export const ColorSelector: React.FC<ColorSelectorProps> = ({
	className,
	disabled,
	selected,
	onClick = () => {},
}) => {
	let colors = ColorArray;

	let badges = colors
		.filter((color) => {
			return color.endsWith("1");
		})
		.map((color, index) => {
			let status: "none" | "selected" | "disabled" = "none";
			if (color === selected) {
				status = "selected";
			}
			if (disabled.includes(color)) {
				status = "disabled";
			}
			return <ColorBadge key={index} status={status} onClick={onClick} color={color}></ColorBadge>;
		});

	return <div className={`${className} flex gap-4`}>{badges}</div>;
};
