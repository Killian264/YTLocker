import React from 'react';
import PropTypes from 'prop-types';
import { PlusButton } from './PlusButton';
import { StatsCard } from './StatsCard';
import { Badge } from './Badge';

export const Card = ({ children, className, ...props }) => {
	return (
		<div className={`${className} bg-secondary py-4 px-6 rounded-md`} {...props}>
			{children[0]}
			<div className="border-b-2 mt-2 mb-3"></div>
			{children.slice(1, children.length)}
		</div>
	);
};

Card.propTypes = {
	className: PropTypes.string,
};