import React, { FC, useState } from 'react';
import { Link, navigate } from '@reach/router';
import {
  Collapse,
  Navbar,
  NavbarToggler,
  Nav,
  NavItem,
  NavLink,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
} from 'reactstrap';
import PathPrefixProps from './types/PathPrefixProps';

const Navigation: FC<PathPrefixProps> = ({ pathPrefix }) => {
  const [isOpen, setIsOpen] = useState(false);
  const toggle = () => setIsOpen(!isOpen);
  return (
    <Navbar className="mb-3" dark color="dark" expand="md" fixed="top">
      <NavbarToggler onClick={toggle} />
      <Link className="pt-0 navbar-brand" to={`${pathPrefix}/new/graph`}>
        Prometheus
      </Link>
      <Collapse isOpen={isOpen} navbar style={{ justifyContent: 'space-between' }}>
        <Nav className="ml-0" navbar>
          <NavItem>
            <NavLink tag={Link} to={`${pathPrefix}/new/alerts`}>
              Alerts
            </NavLink>
          </NavItem>
          <NavItem>
            <NavLink
              tag={Link}
              to={`${pathPrefix}/new/graph`}
              onClick={e => {
                e.preventDefault();
                if (window.location.pathname === `${pathPrefix}/new/graph`) {
                  return;
                }
                navigate(`${pathPrefix}/new/graph${window.location.search}`);
              }}
            >
              Graph
            </NavLink>
          </NavItem>
          <UncontrolledDropdown nav inNavbar>
            <DropdownToggle nav caret>
              Status
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem tag={Link} to={`${pathPrefix}/new/status`}>
                Runtime & Build Information
              </DropdownItem>
              <DropdownItem tag={Link} to={`${pathPrefix}/new/tsdb-status`}>
                TSDB Status
              </DropdownItem>
              <DropdownItem tag={Link} to={`${pathPrefix}/new/flags`}>
                Command-Line Flags
              </DropdownItem>
              <DropdownItem tag={Link} to={`${pathPrefix}/new/config`}>
                Configuration
              </DropdownItem>
              <DropdownItem tag={Link} to={`${pathPrefix}/new/rules`}>
                Rules
              </DropdownItem>
              <DropdownItem tag={Link} to={`${pathPrefix}/new/targets`}>
                Targets
              </DropdownItem>
              <DropdownItem tag={Link} to={`${pathPrefix}/new/service-discovery`}>
                Service Discovery
              </DropdownItem>
            </DropdownMenu>
          </UncontrolledDropdown>
          <NavItem>
            <NavLink href="https://prometheus.io/docs/prometheus/latest/getting_started/">Help</NavLink>
          </NavItem>
          <NavItem>
            <NavLink href={`${pathPrefix}/graph${window.location.search}`}>Classic UI</NavLink>
          </NavItem>
        </Nav>
      </Collapse>
    </Navbar>
  );
};

export default Navigation;
