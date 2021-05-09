import React from 'react';
import PropTypes from 'prop-types';

const colors = {
	primary:   "text-accent-text bg-accent disabled:bg-accent-disabled disabled:text-accent-text-disabled",
	secondary: "text-secondary-text bg-secondary disabled:bg-secondary-disabled disabled:text-secondary-text-disabled"
}


export const Badge = ({ className, color="primary", children, ...props }) => {

	let size = "px-1 text-xs rounded-md"

	return (
		<span
			type="button"
			className={`${className} ${size} ${colors[color]} font-bold select-none`}
			{...props}
		>
			{children}
		</span>
	);
};

Badge.propTypes = {
	color: PropTypes.oneOf(['primary', 'secondary']),
	className: PropTypes.string,
};
