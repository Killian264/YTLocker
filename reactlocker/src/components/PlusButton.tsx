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
	let size = "text-sm rounded-lg inline-block select-none";

	return (
		<button
			disabled={disabled}
			type="button"
			className={`${className} ${size} ${colors[color]}`}
			{...props}
		>
			<span className="text-4xl leading-none font-bold">
				<Plus size={31.5}></Plus>
			</span>
		</button>
	);
};

export interface PlusProps {
	size: number;
}

export const Plus: React.FC<PlusProps> = ({ size }) => {
	return (
		<svg width={size} height={size} fill="none" viewBox="0 0 24 24">
			<path
				stroke="currentColor"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="1.9"
				d="M12 5.75V18.25"
			/>
			<path
				stroke="currentColor"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="1.9"
				d="M18.25 12L5.75 12"
			/>
		</svg>
	);
};
