import React from "react";
import { ColorToColorCSS } from "../shared/colors";
import { Color } from "../shared/types";

export interface ColorBadgeProps {
	className?: string;
	status?: "none" | "disabled" | "selected";
	color: Color;
	onClick?: (color: Color) => void;
}

export const ColorBadge: React.FC<ColorBadgeProps> = ({
	className,
	color,
	status = "none",
	onClick = () => {},
}) => {
	let size = "text-xs leading-none rounded";

	let colorCSS = `bg-${ColorToColorCSS[color]} text-${ColorToColorCSS[color]}`;

	let wrapperCSS = status === "disabled" ? "primary-600" : "yellow-300";

	let badge = (
		<span
			onClick={() => {
				if (status === "none") {
					onClick(color);
				}
			}}
			className={`${className} ${size} ${colorCSS} block w-5 font-bold select-none px-4 py-2`}
		></span>
	);

	if (status === "none") {
		return badge;
	}

	return (
		<div className="flex">
			<span className="line-container">
				<div className={`line-line bg-${wrapperCSS}`}></div>
			</span>
			<div className={`ring-5 ring-${wrapperCSS} rounded`}>
				<div className={`ring-2 ring-primary-800 rounded`}>{badge}</div>
			</div>
		</div>
	);
};
