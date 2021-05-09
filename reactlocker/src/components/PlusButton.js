import React from 'react';
import PropTypes from 'prop-types';

const colors = {
	primary:   "text-accent-text    bg-accent    hover:bg-accent-hover    disabled:bg-accent-disabled    disabled:text-accent-text-disabled",
	secondary: "text-secondary-text bg-secondary hover:bg-secondary-hover disabled:bg-secondary-disabled disabled:text-secondary-text-disabled"
}


export const PlusButton = ({ className, color="primary", disabled, children, ...props }) => {

	let size = "px-2 py-0.5 text-sm rounded-lg inline-block select-none"

	return (
	<button
		disabled={ disabled }
		type="button"
		className={`${className} ${size} ${colors[color]}`}
		{...props}
	>
		<div className="text-4xl leading-none">+</div>
	</button>
	);
};

PlusButton.propTypes = {
	color: PropTypes.oneOf(['primary', 'secondary']),
	className: PropTypes.string,
	disabled: PropTypes.bool,
};
