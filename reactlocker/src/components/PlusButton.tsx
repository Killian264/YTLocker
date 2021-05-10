import React from "react";
import PropTypes from "prop-types";

export interface PlusButtonProps
	extends React.ButtonHTMLAttributes<HTMLButtonElement> {
	className?: string;
	color: keyof typeof colors;
	disabled: boolean;
}

const colors = {
	primary:
		"text-accent-text bg-accent hover:bg-accent-hover disabled:bg-accent-disabled disabled:text-accent-text-disabled",
	secondary:
		"text-secondary-text bg-secondary hover:bg-secondary-hover disabled:bg-secondary-disabled disabled:text-secondary-text-disabled",
};

export const PlusButton: React.FC<PlusButtonProps> = ({
	className,
	color,
	disabled,
	children,
	...props
}) => {
	let size = "px-2 text-sm rounded-lg inline-block select-none";

	return (
		<button
			disabled={disabled}
			type="button"
			className={`${className} ${size} ${colors[color]}`}
			{...props}
		>
			<div className="text-4xl leading-none">+</div>
		</button>
	);
};
