import React from "react";
import { Plus } from "./Svg";

export interface PlusButtonProps
	extends React.ButtonHTMLAttributes<HTMLButtonElement> {
	className?: string;
	color: keyof typeof colors;
}

const colors = {
	primary:
		"text-accent-text bg-accent hover:bg-accent-hover disabled:bg-accent-disabled disabled:text-accent-text-disabled",
	secondary:
		"text-primary-text bg-primary-600 hover:bg-primary-500 disabled:bg-primary-700 disabled:text-primary-text-disabled",
};

export const PlusButton: React.FC<PlusButtonProps> = ({
	className,
	color,
	children,
	...props
}) => {
	let size = "text-sm rounded-lg inline-block select-none";

	return (
		<button
			type="button"
			className={`${className} ${size} ${colors[color]} focus:outline-none`}
			{...props}
		>
			<span className="text-4xl leading-none font-bold">
				<Plus size={31.5}></Plus>
			</span>
		</button>
	);
};
