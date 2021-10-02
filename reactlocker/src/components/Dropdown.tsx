import { Fragment, useState } from "react";
import { Menu, Transition } from "@headlessui/react";
import { Checkmark, ChevronDown } from "./Svg";

export interface DropdownItem {
	title: string;
	value: number;
}

export interface DropdownProps {
	className?: string;
	title: string;
	items: DropdownItem[];
	defaultSelected?: string;
	disabled?: boolean;
	OnItemSelected?: (item: number) => void;
}

export const Dropdown: React.FC<DropdownProps> = ({
	className,
	title,
	items,
	defaultSelected = "",
	disabled = false,
	OnItemSelected = () => {},
}) => {
	const [selected, setSelected] = useState(defaultSelected);

	let dropdownItems = items.map((item, index) => {
		return (
			<MenuItem
				key={index}
				OnClick={() => {
					setSelected(item.title);
					OnItemSelected(item.value);
				}}
				text={item.title}
			></MenuItem>
		);
	});

	return (
		<Menu as="div" className={`${className} relative inline-block text-left`}>
			<div>
				<Menu.Button
					className={`${
						disabled ? "cursor-default" : "hover:bg-primary-500"
					} inline-flex justify-center w-full rounded-md shadow-sm px-4 py-2 bg-primary-600 text-sm font-semibold text-white focus:outline-none`}
				>
					<span className="mt-0.5">{selected === "" ? title : selected}</span>
					<div className="ml-2 -mr-1">
						{selected === "" ? (
							<div className="h-5 w-5 mt-1">
								<ChevronDown></ChevronDown>
							</div>
						) : (
							<Checkmark className="text-green-500" size={24}></Checkmark>
						)}
					</div>
				</Menu.Button>
			</div>

			{disabled || (
				<Transition
					as={Fragment}
					enter="transition ease-out duration-100"
					enterFrom="transform opacity-0 scale-95"
					enterTo="transform opacity-100 scale-100"
					leave="transition ease-in duration-75"
					leaveFrom="transform opacity-100 scale-100"
					leaveTo="transform opacity-0 scale-95"
				>
					<Menu.Items className="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg bg-primary-600 ring-1 ring-black ring-opacity-5 focus:outline-none">
						<div className="py-1">{dropdownItems}</div>
					</Menu.Items>
				</Transition>
			)}
		</Menu>
	);
};

export interface MenuItemProps {
	text: string;
	OnClick: () => void;
}

export const MenuItem: React.FC<MenuItemProps> = ({ text, OnClick }) => {
	const css = "block px-4 py-2 text-sm font-semibold";

	return (
		<Menu.Item>
			{({ active }) => (
				<span
					onClick={OnClick}
					className={`${active ? "bg-primary-500 text-white" : "text-white"} ${css}`}
				>
					{text}
				</span>
			)}
		</Menu.Item>
	);
};
