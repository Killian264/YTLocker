import React from "react";

export interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
	color: keyof typeof colors;
	size: keyof typeof sizes;
	disabled?: boolean;
	loading?: boolean;
}

const sizes = {
	large: "py-2 px-7 text-sm rounded-lg",
	medium: "py-1.5 px-5 text-sm rounded-md",
	small: "py-1 px-3 text-xs rounded-md",
};

const colors = {
	primary:
		"text-accent-text    bg-accent    hover:bg-accent-hover    disabled:bg-accent-disabled    disabled:text-accent-text-disabled",
	secondary:
		"text-primary-text bg-primary-600 hover:bg-primary-500 disabled:bg-primary-700 disabled:text-primary-text-disabled",
};

export const Button: React.FC<ButtonProps> = ({
	className,
	size,
	color,
	disabled = false,
	loading = false,
	children,
	...props
}) => {
	return (
		<button
			disabled={disabled || loading}
			type="button"
			className={`${className} ${sizes[size]}  ${colors[color]} font-bold focus:outline-none`}
			{...props}
		>
			{children}
		</button>
	);
};
