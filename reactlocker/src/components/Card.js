import React from 'react';
import PropTypes from 'prop-types';
import { PlusButton } from './PlusButton';
import { StatsCard } from './StatsCard';
import { Badge } from './Badge';

export const Card = ({ children, ...props }) => {
	return (
		<div className={`bg-secondary py-4 px-6 rounded-md`}>
			<div className="flex justify-between">
				<div>
					<span className=" text-2xl font-bold inline-block align-bottom leading-none">Killian</span>
					<Badge className="ml-1 mt-3" >PRO</Badge>
				</div>
				<div className="flex flex-col justify-end">
					<span className="text-sm leading-none">Joined Mar 28 2020</span>
				</div>
			</div>
			<div className="border-b-2 mt-2 mb-3"></div>
			<div>
				<div className="grid xs:grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 max-w-7xl">
					<StatsCard header={props.header || "Playlists"} count={props.count || "454"} measurement={props.measurement || "total"} {...props} />
					<StatsCard header="Videos" count="357" measurement="total" />
					<StatsCard header="Subscriptions" count="17" measurement="total" />
					<StatsCard header="Updated" count="13" measurement="seconds ago" />
				</div>
			</div>
		</div>
	);
};

Card.propTypes = {};
