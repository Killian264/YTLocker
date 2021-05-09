import React from 'react';
import PropTypes from 'prop-types';


const sizes = {
	large: "py-2 px-6 text-sm rounded-lg",
	medium: "py-1.5 px-5 text-sm rounded-md",
	small: "py-1 px-3 text-xs rounded-md",
}

const colors = {
	primary:   "text-accent-text    bg-accent    hover:bg-accent-hover    disabled:bg-accent-disabled    disabled:text-accent-text-disabled",
	secondary: "text-secondary-text bg-secondary hover:bg-secondary-hover disabled:bg-secondary-disabled disabled:text-secondary-text-disabled"
}


export const Button = ({ className, children, size="large", color="primary", disabled, loading, ...props }) => {
	return (
	<button
		disabled={ disabled || loading }
		type="button"
		className={`${className} ${sizes[size]}  ${colors[color]} font-bold`}
		{...props}
	>
		{children}
	</button>
	);
};

Button.propTypes = {
	size: PropTypes.oneOf(['small', 'medium', 'large']),
	color: PropTypes.oneOf(['primary', 'secondary']),
	className: PropTypes.string,
	disabled: PropTypes.bool,
	loading: PropTypes.bool,
};
