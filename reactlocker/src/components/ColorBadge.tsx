import React from "react";
import { ColorToColorCSS } from "../shared/colors";
import { Color } from "../shared/types";

export interface ColorBadgeProps {
	className?: string;
	color: Color;
}

export const ColorBadge: React.FC<ColorBadgeProps> = ({ className, color }) => {
	let size = "text-xs leading-none rounded";

	let colorCSS = `bg-${ColorToColorCSS[color]} text-${ColorToColorCSS[color]}`;

	return <span className={`${className} ${size} ${colorCSS} font-bold select-none`}>-------</span>;
};
