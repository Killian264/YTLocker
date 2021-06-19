import { Dialog, Transition } from "@headlessui/react";
import { Fragment, useState } from "react";
import { Button } from "./Button";
import { Checkbox } from "./Checkbox";

export interface ModalProps {
	className?: string;
	header: string;
	body: string;
	confirmMessage?: string;
	AcceptClick: () => void;
	RejectClick: () => void;
}

export const Modal: React.FC<ModalProps> = ({
	header,
	body,
	AcceptClick,
	RejectClick,
	confirmMessage = "",
}) => {
	let [isOpen, setIsOpen] = useState(true);
	let [checked, setChecked] = useState(confirmMessage === "");

	const reject = () => {
		setIsOpen(false);
		setTimeout(RejectClick, 300);
	};

	const accept = () => {
		setIsOpen(false);
		setTimeout(AcceptClick, 300);
	};

	const modalCSSReqs = "inline-block w-full overflow-hidden align-middle transition-all transform";

	return (
		<ModalDialogWrapper isOpen={isOpen} CloseClick={reject}>
			<div
				className={`${modalCSSReqs} bg-primary-700 max-w-xs p-6 my-8 text-left rounded-2xl border-2 border-accent`}
			>
				<span className="text-lg font-medium leading-6">{header}</span>
				<p className="mt-2 text-sm">{body}</p>
				{confirmMessage === "" || (
					<Checkbox
						className="mt-2"
						message={confirmMessage}
						checked={checked}
						setChecked={setChecked}
					></Checkbox>
				)}
				<div className="mt-4 flex justify-between">
					<Button onClick={accept} size={"medium"} color={"primary"} disabled={!checked}>
						Confirm
					</Button>
					<Button onClick={reject} className={"ml-2"} size={"medium"} color={"secondary"}>
						Reject
					</Button>
				</div>
			</div>
		</ModalDialogWrapper>
	);
};

export interface ModalDialogWrapperProps {
	isOpen: boolean;
	CloseClick: () => void;
}

export const ModalDialogWrapper: React.FC<ModalDialogWrapperProps> = ({ isOpen, CloseClick, children }) => {
	return (
		<Transition appear show={isOpen} as={Fragment}>
			<Dialog as="div" className="fixed inset-0 z-10 overflow-y-auto" onClose={CloseClick}>
				<div className="min-h-screen px-4 text-center">
					<Transition.Child
						as={Fragment}
						enter="ease-out duration-300"
						enterFrom="opacity-0"
						enterTo="opacity-100"
						leave="ease-in duration-200"
						leaveFrom="opacity-100"
						leaveTo="opacity-0"
					>
						<Dialog.Overlay className="fixed inset-0" />
					</Transition.Child>

					<span className="inline-block h-screen align-middle" aria-hidden="true">
						&#8203;
					</span>
					<Transition.Child
						as={Fragment}
						enter="ease-out duration-300"
						enterFrom="opacity-0 scale-95"
						enterTo="opacity-100 scale-100"
						leave="ease-in duration-200"
						leaveFrom="opacity-100 scale-100"
						leaveTo="opacity-0 scale-95"
					>
						{children}
					</Transition.Child>
				</div>
			</Dialog>
		</Transition>
	);
};
